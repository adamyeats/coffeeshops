package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/db"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/graph"

	"github.com/gin-gonic/gin"
)

// graphqlHandler creates a new GraphQL server
func createGraphqlHandler(db *db.DB) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewSchemaWithResolvers(db))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandler creates a new GraphQL playground
func createPlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
