//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
	"github.com/moyu-x/infinite-synthesis/pkg/gorm"
	"github.com/moyu-x/infinite-synthesis/pkg/server/app"
	"github.com/moyu-x/infinite-synthesis/pkg/server/http/fiber"
)

var pkgProviderSet = wire.NewSet(fiber.NewServer, app.NewApp, gorm.NewGORM)

func initServer(c *config.Bootstrap) *app.App {
	panic(wire.Build(pkgProviderSet))
}
