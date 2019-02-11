package main

import (
	"fmt"
	"os"
	"todo-list-back/config"
	"todo-list-back/database"
	"todo-list-back/logger"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// App application struct
type App struct {
	*echo.Echo
	db *gorm.DB
}

func main() {
	url := config.Conf.Database.GetURL()
	db, err := database.ConnectDB(url)
	if err != nil {
		logger.Err.Errorln(err)
		os.Exit(0)
	}

	app := &App{echo.New(), db}

	app.initRouter()

	defer app.db.Close()

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			context.Set("db", app.db)
			return next(context)
		}
	})

	addr := fmt.Sprintf(":%d", config.Conf.Port)
	// Echo web server start
	if err := app.Start(addr); err != nil {
		logger.Err.Errorln(err)
		app.Logger.Fatal(err)
	}

}
