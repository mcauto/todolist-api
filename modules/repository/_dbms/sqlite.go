package _dbms

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDialector is a sqlite database connection
type SQLiteDialector gorm.Dialector

// NewSQLiteDialector creates a new sqlite dialector
func NewSQLiteDialector() SQLiteDialector {
	return sqlite.Open("todolist.db")
}
