package repository

import (
	"todolist-api/modules/repository/_mysql"

	"go.uber.org/fx"
)

// Modules is a list of all modules
var Modules = fx.Options(fx.Provide(_mysql.NewRepository))
