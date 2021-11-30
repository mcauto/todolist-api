package modules

import (
	"context"
	"fmt"
	"log"
	"todolist-api/modules/config"
	"todolist-api/modules/delivery"
	"todolist-api/modules/repository"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

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

// ToBeInjected 최종 주입될 모듈
var ToBeInjected = fx.Options(
	config.Modules,
	repository.Modules,
	delivery.Modules,
	fx.Invoke(registerHook),
)
