package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$GOPATH/src/todo-list-back")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.UnmarshalKey("app", &Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Test mode")
}

func TestGetURL(t *testing.T) {
	fmt.Println(Conf.Database.GetURL())
}
