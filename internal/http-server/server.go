package httpserver

import (
	"net/http"
	"time"

	"github.com/Paheter1/url-shortener/internal/http-server/handlers"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(address string, timeout, idleTimeout time.Duration) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/shorten", handlers.Shorten)
	return &Server{
		httpServer: &http.Server{
			Addr:        address,
			Handler:     mux,
			ReadTimeout: timeout,
			IdleTimeout: idleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
