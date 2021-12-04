package repository

import (
	"todolist-api/modules/repository/_dbms"

	"go.uber.org/fx"
)

// Modules is a list of all modules
var Modules = fx.Options(fx.Provide(_dbms.NewRepository))
