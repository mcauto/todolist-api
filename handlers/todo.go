package handlers

import (
	"net/http"
	"strconv"
	"todo-list-back/model/todo"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// GetTodoList get todo list
func GetTodoList(context echo.Context) error {
	items, err := todo.FindItem(context.Get("db").(*gorm.DB))
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
func GetTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item := &todo.Item{ID: uint(id)}
	if err := item.Find(context.Get("db").(*gorm.DB)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return context.JSON(http.StatusOK, item)
}

// CreateTodoItem create todo item
func CreateTodoItem(context echo.Context) error {
	item := &todo.Item{}
	if err := context.Bind(item); err != nil {
		return err
	}

	if err := item.Create(context.Get("db").(*gorm.DB)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, item)
}

// UpdateTodoItem update todo item
func UpdateTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item := &todo.Item{ID: uint(id)}
	if err := context.Bind(item); err != nil {
		return err
	}

	if err := item.Update(context.Get("db").(*gorm.DB)); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, item)

}

// DeleteTodoItem delete todo item
func DeleteTodoItem(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item := &todo.Item{ID: uint(id)}

	if err := item.Delete(context.Get("db").(*gorm.DB)); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, item)
}
