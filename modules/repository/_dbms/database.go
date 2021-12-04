package _dbms

import (
	"log"
	"os"
	"time"
	"todolist-api/modules/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// RepositoryImpl is an implementation of Repository
type Repository struct {
	*gorm.DB
}

// NewRepository returns a new instance of Repository
func NewRepository(settings *config.Settings) *Repository {
	var _logger logger.Interface = nil
	if settings.Debug() {
		_logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	db, err := gorm.Open(mysql.Open(settings.DatabaseDSN()), &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{db}
}
