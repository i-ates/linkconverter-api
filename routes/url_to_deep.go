package routes

import (
	"github.com/labstack/echo/v4"
	"linkconverter-api/services"
	"net/http"
)

type UrlToDeepRouterInterface interface {
	UrlToDeep(context echo.Context) error
}

type UrlToDeepRouter struct {
	linkConverterService services.LinkConverterServiceInterface
}

func (urlToDeepRouter *UrlToDeepRouter) UrlToDeep(context echo.Context) error {
	response, err := urlToDeepRouter.linkConverterService.ConvertUrlToDeep(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, nil)
	}

	return context.JSON(http.StatusOK, response)
}

func NewUrlToDeepRouter(linkConverterService services.LinkConverterServiceInterface) UrlToDeepRouterInterface {
	return &UrlToDeepRouter{
		linkConverterService: linkConverterService,
	}
}
