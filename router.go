package main

import (
	"todo-list-back/handlers"

	"github.com/labstack/echo"
)

// Init router init
func (app *App) initRouter() {
	v1 := app.Group("/api/v1")
	{
		v1.GET("/todos", echo.HandlerFunc(handlers.GetTodoList))
		v1.GET("/todos/:id", echo.HandlerFunc(handlers.GetTodoItem))
		v1.POST("/todos", echo.HandlerFunc(handlers.CreateTodoItem))
		v1.PATCH("/todos/:id", echo.HandlerFunc(handlers.UpdateTodoItem))
		v1.DELETE("/todos/:id", echo.HandlerFunc(handlers.DeleteTodoItem))
	}
}
