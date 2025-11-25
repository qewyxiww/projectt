package server

import (
	"log"
	"net/http"
	"time"

	"github.com/qewyxiww/projectt/internal/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Server {

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}
