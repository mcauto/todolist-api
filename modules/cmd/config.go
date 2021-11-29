package cmd

import (
	"fmt"
	"log"
	"todolist-api/modules/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Version: config.AppVersion,
	Short:   fmt.Sprintf("%s Config", config.AppName),
	Run: func(cmd *cobra.Command, args []string) {
		log.Print(config.Json())
	},
}
