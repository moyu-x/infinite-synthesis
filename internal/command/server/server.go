package server

import "context"

func Run(configPath string) {
	ctx := context.Background()
	server := initServer()
	err := server.Start(ctx)
	if err != nil {
		panic(err)
	}
}
