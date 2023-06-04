package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler/transport"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var scopeCtxKey = &contextKey{"scope"}

type contextKey struct {
	name string
}

// ScopeHTTPMiddleware passes the (IAM) scope from the HTTP headers into the context
func ScopeHTTPMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(scopedCtx(r.Context(), r.Header.Get("scope")))
			next.ServeHTTP(w, r)
		})
	}
}

// ScopeWebsocketInit passes the scope information parsed from the websocket init message payload into
// the context, making it available in the resolver
func ScopeWebsocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	return scopedCtx(ctx, initPayload.GetString("scope")), nil
}

// scopedCtx attaches the scope to the context
func scopedCtx(ctx context.Context, scope string) context.Context {
	return context.WithValue(ctx, scopeCtxKey, scope)
}

// GetScopeWithDefault returns the defaultScope if not nil or the token from context
func GetScopeWithDefault(ctx context.Context, defaultScope *string) string {
	// allow override of scope from a string pointer
	if defaultScope != nil && *defaultScope != "" {
		return *defaultScope
	}

	return GetScope(ctx)
}

// GetScope returns the token from context, if provided defaultScope is presented, it will take precedence
func GetScope(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(scopeCtxKey))
}
