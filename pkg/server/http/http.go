package http

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/rs/zerolog"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
	"github.com/moyu-x/infinite-synthesis/third_party/fiberzerolog"
)

type Server struct {
	*fiber.App
	c *config.Bootstrap
	l *zerolog.Logger
}

func NewServer(c *config.Bootstrap, l *zerolog.Logger) *Server {
	return &Server{
		c: c,
		l: l,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.newFiber()

	listenConfig := fiber.ListenConfig{
		GracefulContext: ctx,
	}

	go func() {
		if err := s.App.Listen(s.c.Server.Http.Host+":"+strconv.FormatInt(s.c.Server.Http.Port, 10), listenConfig); err != nil {
			panic("init Server error")
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.App.ShutdownWithContext(ctx); err != nil {
		return errors.New("stop fiber server error")
	}

	return nil
}

func (s *Server) newFiber() {
	fc := fiber.Config{
		JSONEncoder:     sonic.Marshal,
		JSONDecoder:     sonic.Unmarshal,
		AppName:         s.c.Server.Name,
		ReadTimeout:     time.Duration(s.c.Server.Http.Timeout) * time.Second,
		StructValidator: &Validator{},
	}
	s.App = fiber.New(fc)
	s.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: s.l,
	}))
	s.App.Use(pprof.New(pprof.Config{Prefix: "/system/"}))

	s.App.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
}
