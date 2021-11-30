package web

import "go.uber.org/fx"

// Modules is a list of all modules.
var Modules = fx.Options(fx.Provide(NewServer))
