package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/adamyeats/coffeeshops/graphql-server/internal/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a context struct
type Resolver struct{ DB *db.DB }

func NewSchemaWithResolvers(db *db.DB) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{db},
	})
}
