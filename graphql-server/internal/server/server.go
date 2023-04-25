package server

import (
	"context"
	"net/http"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/db"
	"github.com/adamyeats/coffeeshops/graphql-server/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server is the main server struct
type Server struct {
	db     *db.DB
	router *gin.Engine
	server *http.Server
}

// New creates a new server instance
func New(ctx context.Context, config *config.Config) (*Server, error) {
	r := gin.Default()

	// enable CORS
	cnf := cors.DefaultConfig()
	cnf.AllowAllOrigins = true

	r.Use(cors.New(cnf))

	srv := &http.Server{
		Addr:    config.Server.Port,
		Handler: r,
	}

	db, err := db.New(ctx, config.DB)

	if err != nil {
		panic(err)
	}

	// bind routes
	r.POST("/graphql", createGraphqlHandler(db))
	r.GET("/", createPlaygroundHandler())

	return &Server{
		db:     db,
		router: r,
		server: srv,
	}, nil
}

// Serve starts the server
func (s *Server) Serve() error {
	err := s.router.Run()

	if err != nil {
		return err
	}

	return nil
}

// Stop stops the server
func (s *Server) Stop(ctx context.Context) error {
	s.db.Close()
	return s.server.Shutdown(ctx)
}
