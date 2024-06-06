package cmd

import (
	"github.com/moyu-x/infinite-synthesis/internal/command/server"

	"github.com/spf13/cobra"
)

var config string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start web server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(config)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&config, "config", "c", "config.toml", "config file")
}
