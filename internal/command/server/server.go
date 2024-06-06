package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
)

func Run(configPath string) {
	ctx := context.Background()
	conf := config.NewConfig(configPath)

	server := initServer(conf)
	if err := server.Start(ctx); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	// log.Info().Msg("Shutdown Server ...")
	if err := server.Stop(ctx); err != nil {
		panic(err)
	}
}
