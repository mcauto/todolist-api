package domains

import (
	"todolist-api/modules/domains/todo"

	"go.uber.org/fx"
)

// Modules is a list of all modules
var Modules = fx.Options(fx.Provide(todo.NewService))
