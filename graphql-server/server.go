package main

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/xadamy/coffeeshops/graphql-server/graph"
	"github.com/xadamy/coffeeshops/graphql-server/graph/generated"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
