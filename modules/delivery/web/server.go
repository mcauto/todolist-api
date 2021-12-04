package web

import (
	"context"
	"fmt"
	"log"
	"todolist-api/modules/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

// NewServer creates a new server.
func NewServer(settings *config.Settings) *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.RequestID(),
	)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Validator = &RequestValidator{validator: validator.New()}

	return e
}

func registerHook(lifecycle fx.Lifecycle, server *echo.Echo, settings *config.Settings) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := server.Start(settings.BindAddress())
				if err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping Http Server.")
			return server.Shutdown(ctx)
		},
	})
}
