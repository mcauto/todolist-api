package repository

import (
	"todolist-api/modules/repository/_dbms"

	"go.uber.org/fx"
)

// Modules of repository
var Modules = fx.Options(_dbms.Modules)
