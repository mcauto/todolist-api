package database

import (
	"todo-list-back/model/todo"

	"github.com/jinzhu/gorm"
	// gorm에 mysql을 연결하여 사용하기 위한 명시적 import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var url string

func init() {
	url = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
}

// ConnectDB database connection
// TODO: create connection pool by mode (default: debug)
// user:password@/dbname?charset=utf8&parseTime
func ConnectDB(url string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	autoMigrate(db)
	return db, nil
}
func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&todo.Item{})
}
