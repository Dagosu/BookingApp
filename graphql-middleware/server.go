package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/example/starwars/generated"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	ghc "github.com/Dagosu/BookingApp/gohelpers/config"
	"github.com/Dagosu/BookingApp/gohelpers/configload"
	"github.com/Dagosu/BookingApp/graphql-middleware/config"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph"
	"github.com/Dagosu/BookingApp/graphql-middleware/graph/auth"
	"github.com/Dagosu/BookingApp/graphql-middleware/middleware"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	configload.Hydrate(&config.C)
	configload.Hydrate(&ghc.C)

	s, err := graph.NewGraphQLServer()
	if err != nil {
		log.Fatalln("Cannot create new GraphQL server", err)
	}
	defer s.Close()

	router := chi.NewRouter()

	// GraphQL entry point
	router.Group(func(router chi.Router) {
		router.Use(auth.HTTPMiddleware())
		router.Use(middleware.ScopeHTTPMiddleware())

		allowedOrigins := []string{
			fmt.Sprintf("http://localhost:%s", config.C.Port),
			"electron://altair",
			"*", // TODO remove wildcard origin
		}
		// Add CORS middleware around every request
		// See https://github.com/rs/cors for full option listing
		router.Use(cors.New(cors.Options{
			AllowedOrigins:   allowedOrigins,
			AllowCredentials: true,
			Debug:            false,
			AllowedHeaders:   []string{"*"},
		}).Handler)

		router.Handle("/query", router)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.C.Port)
	log.Fatal(http.ListenAndServe(":"+config.C.Port, router))
}

func newGraphQLHandler(s *graph.Server) *handler.Server {
	conf := addDirectives(generated.Config{Resolvers: s})
	es := generated.NewExecutableSchema(conf)
	srv := handler.New(es)

	// limit query complexity
	srv.Use(extension.FixedComplexityLimit(config.C.QueryComplexityLimit))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 15 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true // TODO remove wildcard origin
				return r.Host == "localhost" || r.Host == "electron://altair"
			},
			ReadBufferSize:    1024,
			WriteBufferSize:   1024,
			EnableCompression: true,
		},
		InitFunc: websocketInitChain,
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	// do we need file upload? not yet.
	// srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	if config.C.Development {
		// display panic messages when runnin local
	} else {
		srv.SetRecoverFunc(recoverFunc)
	}

	return srv
}

// addDirectives adds custom directives declared in the schema
// see https://gqlgen.com/reference/directives/
func addDirectives(c generated.Config) generated.Config {
	c.Directives.ResourceMdmResource = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		return next(ctx)
	}
	c.Directives.KPIResourceKpi = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		return next(ctx)
	}
	return c
}

func recoverFunc(ctx context.Context, err interface{}) error {
	log.Printf("Panic: %v", err)
	// notify bug tracker here...
	return errors.New("Internal server error")
}

func websocketInitChain(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	// fetch access token
	chainCtx, err := auth.WebsocketInit(ctx, initPayload)
	if err != nil {
		return ctx, err
	}

	// fetch scope
	return middleware.ScopeWebsocketInit(chainCtx, initPayload)
}
