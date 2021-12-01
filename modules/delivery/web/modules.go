package web

import (
	"todolist-api/modules/delivery/web/v1/todo"

	"go.uber.org/fx"
)

// Modules is a list of all modules.
var Modules = fx.Options(
	fx.Provide(NewServer),
	fx.Provide(todo.NewHandler),
	fx.Invoke(registerHook, todo.BindRoutes),
)
