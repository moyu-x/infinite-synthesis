package app

import (
	"context"

	"github.com/moyu-x/infinite-synthesis/pkg/server/http/fiber"
)

type App struct {
	h *fiber.Server
}

func NewApp(h *fiber.Server) *App {
	return &App{
		h: h,
	}
}

func (a *App) Start(ctx context.Context) error {
	if err := a.h.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.h.Stop(ctx); err != nil {
		return err
	}
	return nil
}
