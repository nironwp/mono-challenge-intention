package graphserver

import (
	"intent-service/domain/entities"
	"intent-service/graph"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GraphServer struct {
	Port    string
	Service entities.IntentService
}

func (gs *GraphServer) Start() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: gs.Service,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", gs.Port)
	log.Fatal(http.ListenAndServe(":"+gs.Port, nil))
}

func NewGraphServer(port string, service entities.IntentService) *GraphServer {
	return &GraphServer{
		Port:    port,
		Service: service,
	}
}
