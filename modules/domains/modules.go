package domains

import (
	"todolist-api/modules/domains/todo"

	"go.uber.org/fx"
)

// Modules
var Modules = fx.Options(fx.Provide(todo.NewService))
