package main

import (
	"fmt"
	"net/http"
	"os"
	"todo-list-back/config"
	"todo-list-back/database"
	"todo-list-back/logger"
	todoHttp "todo-list-back/todo/delivery/http"
	"todo-list-back/todo/repository"
	"todo-list-back/todo/usecase"

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

	todoRepository := repository.NewGormRepository(app.db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHttp.NewTodoHandler(e, todoUsecase)

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
