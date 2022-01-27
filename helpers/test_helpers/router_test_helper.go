package test_helpers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

func SetupEcho(query_string_params string) (*httptest.ResponseRecorder, echo.Context) {
	var path = query_string_params
	if path == "" {
		path = "/"
	}

	echo := echo.New()
	request := httptest.NewRequest(http.MethodGet, path, nil)
	recorder := httptest.NewRecorder()
	context := echo.NewContext(request, recorder)

	return recorder, context
}
