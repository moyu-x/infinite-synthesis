package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate http interface use protobuf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
