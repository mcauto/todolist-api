package modules

import (
	"todolist-api/modules/config"

	"go.uber.org/fx"
)

// ToBeInjected 최종 주입될 모듈
var ToBeInjected = fx.Options(
	config.Modules,
)
