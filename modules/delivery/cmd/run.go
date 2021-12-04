package cmd

import (
	"fmt"
	"log"
	"todolist-api/modules/config"
	"todolist-api/modules/delivery/web"
	"todolist-api/modules/domains"
	"todolist-api/modules/domains/todo"
	"todolist-api/modules/repository/_dbms"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Version: config.AppVersion,
	Short:   fmt.Sprintf("%s Run", config.AppName),
	Run: func(cmd *cobra.Command, args []string) {
		f := func(repo *_dbms.Repository) {
			if err := repo.AutoMigrate(&todo.Item{}); err != nil {
				log.Fatal(err)
			}
		}
		modules := fx.Options(
			config.Modules,
			domains.Modules,
			web.Modules,
			_dbms.Modules,
			fx.Invoke(f),
		)
		fx.New(modules).Run()
	},
}
