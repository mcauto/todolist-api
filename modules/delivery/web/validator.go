package web

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// RequestValidator HTTP 요청에 대한 validator
// https://echo.labstack.com/guide/request/#validate-data
type RequestValidator struct {
	validator *validator.Validate
}

// NewRequestValidator validator 생성자
func NewRequestValidator(validator *validator.Validate) *RequestValidator {
	return &RequestValidator{validator: validator}
}

// Validate 유효성 검사 함수
func (v *RequestValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
