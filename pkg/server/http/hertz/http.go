package hertz

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
)

type Server struct {
	hertz *server.Hertz
	c     *config.Bootstrap
}

func NewServer(c *config.Bootstrap) *Server {
	return &Server{
		c: c,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.hertz = server.Default()
	go func() {
		s.hertz.Spin()
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	_ = s.hertz.Shutdown(ctx)
	return nil
}
