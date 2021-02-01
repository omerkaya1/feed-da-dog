package internal

import (
	"context"
	"fmt"
)

type logger interface {
	Println(args ...interface{})
	SetPrefix(prefix string)
}

type Server struct {
	s Store
	logger
	errChan chan error
}

func NewServer(s Store, l logger, errChan chan error) *Server {
	return &Server{s, l, errChan}
}

func (s *Server) Start(ctx context.Context) error {
	fmt.Println("Started...waiting for the context cancellation")
	<-ctx.Done()
	fmt.Println("Done! Exiting...")
	return nil
}
