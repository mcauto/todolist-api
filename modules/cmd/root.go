package cmd

import (
	"fmt"
	"os"
	"todolist-api/modules/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     config.AppName,
	Version: config.AppVersion,
	Short:   fmt.Sprintf("%s Config", config.AppName),
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(runCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
