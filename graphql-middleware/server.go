package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Dagosu/BookingApp/graphql-middleware/config"
	graph "github.com/Dagosu/BookingApp/graphql-middleware/graph"
	generated "github.com/Dagosu/BookingApp/graphql-middleware/graph/generated"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	s, err := graph.NewGraphQLServer()
	if err != nil {
		log.Fatalln("Cannot create new GraphQL server", err)
	}
	defer s.Close()

	router := chi.NewRouter()
	srv := newGraphQLHandler(s)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// GraphQL entry point
	router.Group(func(router chi.Router) {
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
		router.Use(func(next http.Handler) http.Handler {
			return middleware(srv)
		})

		router.Handle("/query", router)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func newGraphQLHandler(s *graph.Server) *handler.Server {
	conf := addDirectives(generated.Config{Resolvers: s})
	es := generated.NewExecutableSchema(conf)
	srv := handler.New(es)

	// limit query complexity
	srv.Use(extension.FixedComplexityLimit(4000))

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
	return c
}

func recoverFunc(ctx context.Context, err interface{}) error {
	log.Printf("Panic: %v", err)
	// notify bug tracker here...
	return errors.New("Internal server error")
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
