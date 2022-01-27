// Package classification LinkConverterApi
//
// Documentation for Link Converter Api
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
// Produces:
// - application/json
// swagger:meta
package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"linkconverter-api/helpers"
	"linkconverter-api/libs/logging"
	middleware2 "linkconverter-api/middleware"
	"linkconverter-api/routes"
)

func createLogger() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Sampling = nil
	logConfig := logging.LoggerConfig{
		Config: zapConfig,
		ContextFieldFunc: func(ctx context.Context) []zap.Field {
			var fields []zap.Field

			if ctxRequestPath, ok := ctx.Value(helpers.RequestPath).(string); ok {
				fields = append(fields, zap.String(helpers.RequestPath, ctxRequestPath))
			}
			if ctxEnvironment, ok := ctx.Value(helpers.Environment_Const).(string); ok {
				fields = append(fields, zap.String(helpers.Environment_Const, ctxEnvironment))
			}
			return fields
		}}

	logging.GetLogger(logConfig)
}
func init() {
	createLogger()
}

func main() {
	container := BuildContainer()
	api := SetupRouter(*container)
	api.HideBanner = true

	err := container.Invoke(func(
		config helpers.Config,
	) {
		_ = api.Start(fmt.Sprintf(":%s", config.Port))
	})

	if err != nil {
		panic(err)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()

	_ = container.Provide(helpers.NewConfig)
	_ = container.Provide(routes.NewStatusRouter)

	return container
}

func SetupRouter(container dig.Container) *echo.Echo {

	var api = echo.New()

	api.File("/favicon.ico", "")

	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			middleware2.WithDefaultFieldsForLogging(c)
			return next(c)
		}
	})

	err := container.Invoke(func(
		statusRouter routes.StatusRouterInterface,
	) {
		api.GET("/status", statusRouter.Status)
	})

	api.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))

	if err != nil {
		panic(err)
	}

	return api
}
