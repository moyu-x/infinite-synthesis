package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/moyu-x/infinite-synthesis/pkg/config"
	"github.com/moyu-x/infinite-synthesis/pkg/logger"
)

func Run(configPath string) {
	ctx := context.Background()
	conf := config.NewConfig(configPath)
	logger.NewLogger(conf)

	server := initServer(conf)
	if err := server.Start(ctx); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Info().Msg("Shutdown Server ...")
	if err := server.Stop(ctx); err != nil {
		panic(err)
	}
	log.Info().Msg("Shutdown Server Complete.")
}
