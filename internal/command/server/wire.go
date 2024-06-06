//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/moyu-x/infinite-synthesis/pkg/server/http"
)

var pkgProviderSet = wire.NewSet(http.NewServer)

func initServer() *http.Server {
	panic(wire.Build(pkgProviderSet))
}
