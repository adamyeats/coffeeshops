package main

import (
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"

	"go.elastic.co/apm/module/apmelasticsearch/v2"
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

	// init with elastic apm integration
	cfg := elasticsearch.Config{
		Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
	}

	_, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	r.Use(apmgin.Middleware(r))
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
