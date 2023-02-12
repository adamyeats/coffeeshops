package server

import (
	"context"
	"log"
	"net/http"

	"github.com/adamyeats/coffeeshops/graphql-server/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Server is the main server struct
type Server struct {
	db     *db.DB
	router *gin.Engine
	server *http.Server
}

// New creates a new server instance
func New(ctx context.Context, port string) (*Server, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	db, err := db.New(ctx)

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
