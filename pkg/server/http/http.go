package http

import (
	"context"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
