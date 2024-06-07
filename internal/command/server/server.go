package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
	"github.com/moyu-x/infinite-synthesis/pkg/log"
)

func Run(configPath string) {
	ctx := context.Background()
	conf := config.NewConfig(configPath)
	logger := log.NewLogger(conf)

	server := initServer(conf)
	if err := server.Start(ctx); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.Info().Msg("Shutdown Server ...")
	if err := server.Stop(ctx); err != nil {
		panic(err)
	}
	logger.Info().Msg("Shutdown Server Complete.")
}
