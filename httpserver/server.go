package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
	notify          chan error
}

const (
	_defaultTimeOut = 3 * time.Second
	_defaultAddr    = ":8000"
)

func NewServer(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		shutdownTimeout: _defaultTimeOut,
		notify:          make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	go func() {
		s.start()
	}()
	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
