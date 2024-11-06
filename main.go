package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/siyaramsujan/graphql-api/graph"
	"github.com/siyaramsujan/graphql-api/server_lib/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	
  if port == "" {
		port = defaultPort
	}
    
  router := chi.NewRouter()

  newResolver := graph.NewServiceResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: newResolver}))

  router.Handle("/", playground.Handler("GraphQL playground", "/query"))
    
  router.Group(func(r chi.Router) {
     r.Use(middleware.AuthMiddleware())
     r.Handle("/query", srv)
  })

	log.Fatal(http.ListenAndServe(":"+port, router))
}
