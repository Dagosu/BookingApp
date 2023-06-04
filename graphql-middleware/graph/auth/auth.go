package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var tknCtxKey = &contextKey{"access_token"}

// var scopeCtxKey = &contextKey{"scope"}

type contextKey struct {
	name string
}

// HTTPMiddleware passes the authorization token from the HTTP headers into the context
func HTTPMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tkn := r.Header.Get("authorization")
			// Allow unauthenticated users in
			if tkn == "" {
				next.ServeHTTP(w, r)

				return
			}

			r = r.WithContext(authCtx(r.Context(), tkn))
			next.ServeHTTP(w, r)
		})
	}
}

// WebsocketInit passes the authorization information parsed from the websocket init message payload into
// the context, making it available in the resolver
func WebsocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	return authCtx(ctx, initPayload.Authorization()), nil
}

// streamInterceptor returns a client interceptor to authenticate stream RPC
func streamInterceptor() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		log.Printf("--> stream interceptor: %s", method)

		return streamer(attachAccessToken(ctx), desc, cc, method, opts...)
	}
}

// unaryInterceptor returns a client interceptor to authenticate unary RPC
func unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		return invoker(attachAccessToken(ctx), method, req, reply, cc, opts...)
	}
}

// attachAccessToken attaches interceptor token to passed context
func attachAccessToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", GetAuthTkn(ctx))
}

// authCtx attaches the access token to the given context.
func authCtx(ctx context.Context, tkn string) context.Context {
	return context.WithValue(ctx, tknCtxKey, tkn)
}

// GetAuthTkn returns the token from context. Useful for debugging.
func GetAuthTkn(ctx context.Context) string {
	return fmt.Sprintf("%v", ctx.Value(tknCtxKey))
}

// NewAuthContext returns a pristine context with the authorization value attached
// todo: also keep tracing information
func NewAuthContext(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), tknCtxKey, GetAuthTkn(ctx))
}

// WithInterceptors return dial options to be attached to gRPC clients for auth interceptors
func WithInterceptors() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				unaryInterceptor(),
				otelgrpc.UnaryClientInterceptor(),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				streamInterceptor(),
				otelgrpc.StreamClientInterceptor(),
			),
		),
	}
}
