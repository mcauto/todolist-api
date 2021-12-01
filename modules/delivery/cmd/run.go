package cmd

import (
	"fmt"
	"todolist-api/modules/config"
	"todolist-api/modules/delivery/web"
	"todolist-api/modules/domains"
	"todolist-api/modules/domains/todo"
	"todolist-api/modules/repository"
	"todolist-api/modules/repository/_mysql"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Version: config.AppVersion,
	Short:   fmt.Sprintf("%s Run", config.AppName),
	Run: func(cmd *cobra.Command, args []string) {
		f := func(repo *_mysql.Repository) {
			repo.AutoMigrate(&todo.Item{})
		}
		modules := fx.Options(
			config.Modules,
			repository.Modules,
			domains.Modules,
			web.Modules,
			fx.Invoke(f),
		)
		fx.New(modules).Run()
	},
}
