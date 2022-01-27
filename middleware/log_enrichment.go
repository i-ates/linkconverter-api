package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"linkconverter-api/helpers"
	"linkconverter-api/libs/runtimeenvironment"
)

func WithDefaultFieldsForLogging(c echo.Context) {
	path := c.Path()
	if len(c.Request().URL.RawQuery) > 0 {
		path = c.Path() + "?" + c.Request().URL.RawQuery
	}
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, helpers.RequestPath, path)
	ctx = context.WithValue(ctx, helpers.Environment_Const, runtimeenvironment.Environment)
	c.SetRequest(c.Request().WithContext(ctx))
}
