package httpserver

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(address string, timeout, idleTimeout time.Duration) *Server {
	return &Server{}
}
