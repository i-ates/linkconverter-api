package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"linkconverter-api/helpers"
	"linkconverter-api/libs/runtimeenvironment"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceWithDefaultFieldsForLogging(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/synonym?q=laptop")
	WithDefaultFieldsForLogging(c)

	assert.Equal(t, "/synonym?q=laptop", c.Request().Context().Value(helpers.RequestPath))
	assert.Equal(t, runtimeenvironment.Environment, c.Request().Context().Value(helpers.Environment_Const))

}
