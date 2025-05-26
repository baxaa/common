package httpserver

import "time"

type Option func(s *Server)

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

func WithAddr(port string) Option {
	return func(s *Server) {
		s.server.Addr = port
	}
}
