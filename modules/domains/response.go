package domains

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// JSONResponse 응답 값을 JSON으로 반환합니다.
//
// body가 없을 경우 status code의 message를 전달합니다.
func JSONResponse(
	context echo.Context,
	status int,
	body interface{},
) error {
	message := http.StatusText(status)
	if body == nil {
		body = map[string]interface{}{"message": message}
	}
	return context.JSON(status, body)
}
