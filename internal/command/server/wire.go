//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
	"github.com/moyu-x/infinite-synthesis/pkg/log"
	"github.com/moyu-x/infinite-synthesis/pkg/server/app"
	"github.com/moyu-x/infinite-synthesis/pkg/server/http"
)

var pkgProviderSet = wire.NewSet(http.NewServer, app.NewApp, log.NewLogger)

func initServer(c *config.Bootstrap) *app.App {
	panic(wire.Build(pkgProviderSet))
}
