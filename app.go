package main

import (
	"fmt"
	"net/http"
	"os"
	"todo-list-back/config"
	"todo-list-back/database"
	"todo-list-back/logger"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// App application struct
type App struct {
	*echo.Echo
	db *gorm.DB
}

var e *echo.Echo

func init() {
	e = echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
}

func main() {
	url := config.Conf.Database.GetURL()
	db, err := database.ConnectDB(url)
	if err != nil {
		logger.Err.Errorln(err)
		os.Exit(0)
	}

	app := &App{e, db}

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
