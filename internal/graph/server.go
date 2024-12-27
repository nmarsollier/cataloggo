package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nmarsollier/cataloggo/internal/engine/env"
	"github.com/nmarsollier/cataloggo/internal/engine/log"
	"github.com/nmarsollier/cataloggo/internal/graph/model"
	graph "github.com/nmarsollier/cataloggo/internal/graph/schema"
)

func Start(logger log.LogRusEntry) {
	port := env.Get().GqlPort
	srv := handler.NewDefaultServer(model.NewExecutableSchema(model.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("GraphQL playground on port : ", port)
	logger.Error(http.ListenAndServe(fmt.Sprintf(":%d", env.Get().GqlPort), nil))
}