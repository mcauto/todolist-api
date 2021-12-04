package _dbms

import (
	"log"
	"todolist-api/modules/config"

	"go.uber.org/fx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Repository is an implementation of Repository
type Repository struct {
	*gorm.DB
}

// NewRepository returns a new instance of Repository
func NewRepository(
	settings *config.Settings,
	sqliteDialector SQLiteDialector,
	mysqlDialector MySQLDialector,
	_logger logger.Interface,
) *Repository {
	var dialector gorm.Dialector = mysqlDialector
	if settings.Debug() {
		dialector = sqliteDialector
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{db}
}

// Modules returns the module definition
var Modules = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewMySQLDialector),
	fx.Provide(NewSQLiteDialector),
	fx.Provide(NewRepository),
)
