package config

import (
	"encoding/json"
	"log"

	env "github.com/Netflix/go-env"
	"go.uber.org/fx"
)

const (
	// AppVersion version
	AppVersion = "21.1.0"
	// AppName name
	AppName = "todolist-api"
)

// Settings 설정 값
type Settings struct {
	App struct {
		Version     string `env:"-,default=1.0.0" json:"version"`
		Name        string `env:"-,default=todolist-api" json:"name"`
		Port        int    `env:"PORT,default=5000" json:"port"`
		Environment string `env:"ENVIRONMENT,default=development" json:"environment"` // development, stage, production
	} `json:"app"`
	Database struct {
		User     string `env:"DATABASE_USER,required=true" json:"user"`
		Password string `env:"DATABASE_PASSWORD,required=true" json:"password"`
		Name     string `env:"DATABASE_NAME,default=todolist" json:"name"`
		Host     string `env:"DATABASE_HOST,default=localhost" json:"host"`
		Port     int    `env:"DATABASE_PORT,default=3306" json:"port"`
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

// JSON 설정 값 출력
func JSON() string {
	settings := NewSettings()
	jsonBytes, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}

// Modules config 모듈
var Modules = fx.Options(fx.Provide(NewSettings))
