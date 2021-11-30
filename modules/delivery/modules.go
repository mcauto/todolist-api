package delivery

import (
	"todolist-api/modules/delivery/web"

	"go.uber.org/fx"
)

// Modules is a list of all modules.
var Modules = fx.Options(web.Modules)
