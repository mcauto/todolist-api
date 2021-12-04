package _dbms

import (
	"todolist-api/modules/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDialector is a dialector for MySQL
type MySQLDialector gorm.Dialector

// NewMySQLDialector returns a new instance of MySQLDialector
func NewMySQLDialector(
	settings *config.Settings,
) MySQLDialector {
	return mysql.Open(settings.DatabaseDSN())
}
