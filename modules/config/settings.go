package config

import (
	"encoding/json"
	"log"

	env "github.com/Netflix/go-env"
	"go.uber.org/fx"
)

const (
	// App version
	AppVersion = "1.0.0"
	// App name
	AppName = "todolist-api"
)

// Settings 설정 값
type Settings struct {
	App struct {
		Version string `env:"-,default=1.0.0" json:"version"`
		Name    string `env:"-,default=todolist-api" json:"name"`
		Port    int    `env:"PORT" json:"port"`
	} `json:"app"`
	Database struct {
		DSN string `env:"DATABASE_DSN,required=true" json:"dsn"`
	} `json:"database"`
	Extras env.EnvSet `json:"-"`
}

// NewSettings 설정 값 생성
func NewSettings() *Settings {
	var settings Settings
	extras, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal(err)
	}
	settings.Extras = extras

	settings.App.Version = AppVersion
	settings.App.Name = AppName
	return &settings
}

// Json
func Json() string {
	settings := NewSettings()
	jsonBytes, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}

// Modules config 모듈
var Modules = fx.Options(fx.Provide(NewSettings))
