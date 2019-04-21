package http

import (
	"net/http"
	"strconv"

	"todo-list-back/models"
	"todo-list-back/todo"

	"github.com/labstack/echo"
)

// TodoHandler todo http handler
type TodoHandler struct {
	usecase todo.Usecase
}

// NewTodoHandler echo에 http handler 등록
func NewTodoHandler(e *echo.Echo, usecase todo.Usecase) {
	handler := &TodoHandler{
		usecase: usecase,
	}
	v1 := e.Group("/api/v1")
	{
		v1.GET("/todos", echo.HandlerFunc(handler.GetTodoList))
		v1.GET("/todos/:id", echo.HandlerFunc(handler.GetTodoItem))
		v1.POST("/todos", echo.HandlerFunc(handler.CreateTodoItem))
		v1.PATCH("/todos/:id", echo.HandlerFunc(handler.UpdateTodoItem))
		v1.DELETE("/todos/:id", echo.HandlerFunc(handler.DeleteTodoItem))
	}

}

// GetTodoList get todo list
func (handler *TodoHandler) GetTodoList(context echo.Context) error {
	items, err := handler.usecase.FindAll()
	if err != nil {
		return err
	}
	var code int
	if len(items) > 0 {
		code = http.StatusOK
	} else {
		code = http.StatusNoContent
	}
	return context.JSON(code, items)
}

// GetTodoItem get todo item
func (handler *TodoHandler) GetTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item, err := handler.usecase.FindByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return context.JSON(http.StatusOK, item)
}

// CreateTodoItem create todo item
func (handler *TodoHandler) CreateTodoItem(context echo.Context) error {
	item := &models.TodoItem{}
	if err := context.Bind(item); err != nil {
		return err
	}

	err := handler.usecase.Create(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, item)
}

// UpdateTodoItem update todo item
func (handler *TodoHandler) UpdateTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item := &models.TodoItem{ID: uint(id)}
	if err := context.Bind(item); err != nil {
		return err
	}

	err = handler.usecase.Update(item)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, item)

}

// DeleteTodoItem delete todo item
func (handler *TodoHandler) DeleteTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item := &models.TodoItem{ID: uint(id)}

	err = handler.usecase.Delete(item)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, item)
}
