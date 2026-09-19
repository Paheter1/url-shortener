package httpserver

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(address string, timeout, idleTimeout time.Duration) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:        address,
			ReadTimeout: timeout,
			IdleTimeout: idleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
