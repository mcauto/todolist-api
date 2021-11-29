package cmd

import (
	"fmt"
	"todolist-api/modules"
	"todolist-api/modules/config"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Version: config.AppVersion,
	Short:   fmt.Sprintf("%s Run", config.AppName),
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(modules.ToBeInjected).Run()
	},
}
