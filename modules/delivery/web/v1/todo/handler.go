package todo

import (
	"net/http"
	"strconv"
	"todolist-api/modules/domains"
	"todolist-api/modules/domains/todo"

	"github.com/labstack/echo/v4"
)

// Handler todo handler
type Handler struct {
	todo.Service
}

// NewHandler 생성자
func NewHandler(service todo.Service) *Handler {
	return &Handler{service}
}

// BindRoutes todo bind routes
func BindRoutes(server *echo.Echo, handler *Handler) {
	group := server.Group("/todos")
	group.GET("", handler.GetAll)
	group.GET("/:id", handler.Get)
	group.POST("", handler.Post)
	group.PATCH("/:id", handler.Patch)
	group.DELETE("/:id", handler.Delete)
}

// GetAll todo get all
// @ID todo-get-all
// @Tags todo
// @Summary todo get all
// @Description todo get all
// @Router /todos [get]
// @Param _ query Params true "params"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} ManyItemResponse
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) GetAll(c echo.Context) error {
	var param Params
	if err := c.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&param); err != nil {
		return err
	}
	page, limit := param.Pagination()
	items, err := h.FetchAll(page, limit)
	if err != nil {
		return domains.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := ManyItemResponse{
		Todos: items,
		Page:  page + 1,
		Limit: limit,
	}
	return domains.JSONResponse(c, http.StatusOK, response)
}

// Get todo get
// @ID todo-get
// @Tags todo
// @Summary todo get
// @Description todo get
// @Router /todos/{id} [get]
// @Param id path uint64 true "id"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} todo.Item
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Get(c echo.Context) (err error) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item, err := h.Fetch(uint64(id))
	if err != nil {
		return domains.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	return domains.JSONResponse(c, http.StatusOK, item)
}

// Post todo post
// @ID todo-post
// @Tags todo
// @Summary todo post
// @Description todo post
// @Router /todos [post]
// @Param body body Body true "body"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} todo.Item
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Post(c echo.Context) error {
	item := &todo.Item{}
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(item); err != nil {
		return err
	}

	affected, err := h.Insert(item)
	if err != nil {
		return domains.JSONResponse(c, http.StatusInternalServerError, err.Error())
	}
	if affected == 0 {
		return domains.JSONResponse(c, http.StatusConflict, nil)
	}
	return domains.JSONResponse(c, http.StatusOK, item)
}

// Patch todo patch
// @ID todo-patch
// @Tags todo
// @Summary todo patch
// @Description todo patch
// @Router /todos/{id} [patch]
// @Param id path uint64 true "id"
// @Param body body Body true "body"
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} interface{}
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Patch(c echo.Context) error {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	body := &Body{}
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(body); err != nil {
		return err
	}
	updated, err := h.Service.Update(uint64(id), body.Title)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if updated == nil {
		return domains.JSONResponse(c, http.StatusNotFound, nil)
	}
	return domains.JSONResponse(c, http.StatusOK, updated)
}

// Delete todo delete
// @ID todo-delete
// @Tags todo
// @Summary todo delete
// @Description todo delete
// @Router /todos/{id} [delete]
// @Param id path uint64 true "id"
// @Security ApiKeyAuth
// @Produce json
// @Success 204 ""
// @Failure 400 {object} interface{}
// @Failure 500 {object} interface{}
func (h Handler) Delete(c echo.Context) error {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	affected, err := h.Service.Delete(uint64(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if affected == 0 {
		return domains.JSONResponse(c, http.StatusNotFound, nil)
	}

	return domains.JSONResponse(c, http.StatusNoContent, nil)
}
