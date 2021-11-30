package _mysql

import (
	"log"
	"todolist-api/modules/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repository is an interface for the repository
type Repository interface{}

// RepositoryImpl is an implementation of Repository
type RepositoryImpl struct {
	*gorm.DB
}

// NewRepository returns a new instance of Repository
func NewRepository(settings *config.Settings) Repository {
	db, err := gorm.Open(mysql.Open(settings.DatabaseDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &RepositoryImpl{db}
}
