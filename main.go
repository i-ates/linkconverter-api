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
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"linkconverter-api/helpers"
	"linkconverter-api/routes"
	"linkconverter-api/services"
)

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
	_ = container.Provide(routes.NewDeepToUrlRouter)
	_ = container.Provide(routes.NewUrlToDeepRouter)
	_ = container.Provide(services.NewLinkConverterService)

	return container
}

func SetupRouter(container dig.Container) *echo.Echo {

	var api = echo.New()

	api.File("/favicon.ico", "")

	err := container.Invoke(func(
		statusRouter routes.StatusRouterInterface,
		urlToDeepRouter routes.UrlToDeepRouterInterface,
		deepToUrlRouter routes.DeepToUrlRouterInterface,
	) {
		api.GET("/status", statusRouter.Status)
		api.POST("/urltodeep", urlToDeepRouter.UrlToDeep)
		api.POST("deeptourl", deepToUrlRouter.DeepToUrl)
	})

	api.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))

	if err != nil {
		panic(err)
	}

	return api
}
