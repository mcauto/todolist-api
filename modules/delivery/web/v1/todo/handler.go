package todo

import (
	"net/http"
	"todolist-api/modules/domains"
	"todolist-api/modules/domains/todo"

	"github.com/labstack/echo/v4"
)

// Handler
type Handler struct {
	todo.Service
}

// NewHandler 생성자
func NewHandler(service todo.Service) *Handler {
	return &Handler{service}
}

// BindRoutes
func BindRoutes(server *echo.Echo, handler *Handler) {
	group := server.Group("/todos")
	group.GET("", handler.GetAll)
	group.GET("/:id", handler.Get)
	group.POST("", handler.Post)
	group.PATCH("", handler.Patch)
	group.DELETE("", handler.Delete)
}

// GetAll todo get all
// @ID todo-get-all
// @Tags todo
// @Summary todo get all
// @Description todo get all
// @Router /todos [get]
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} []Item
// @Success 204 ""
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) GetAll(c echo.Context) error {
	items, err := h.FetchAll(5, 10)
	if err != nil {
		return domains.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return domains.JSONResponse(c, http.StatusOK, items)
}

// Get todo get
// @ID todo-get
// @Tags todo
// @Summary todo get
// @Description todo get
// @Router /todos/:id [get]
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} interface{}
// @Success 204 ""
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Get(c echo.Context) error {

	return nil
}

// Post todo post
// @ID todo-post
// @Tags todo
// @Summary todo post
// @Description todo post
// @Router /todos [post]
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} interface{}
// @Success 204 ""
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Post(c echo.Context) error {
	return nil
}

// Patch todo patch
// @ID todo-patch
// @Tags todo
// @Summary todo patch
// @Description todo patch
// @Router /todos [patch]
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} interface{}
// @Success 204 ""
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Patch(c echo.Context) error {
	return nil
}

// Delete todo delete
// @ID todo-delete
// @Tags todo
// @Summary todo delete
// @Description todo delete
// @Router /todos [delete]
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} interface{}
// @Success 204 ""
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Delete(c echo.Context) error {
	return nil
}
