package routes

import (
	"github.com/labstack/echo/v4"
	"linkconverter-api/models/requests"
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
	urlRequestModel := requests.UrlRequestModel{}

	err := context.Bind(&urlRequestModel)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	urlToDeepResponseModel, err := urlToDeepRouter.linkConverterService.ConvertUrlToDeep(urlRequestModel)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, urlToDeepResponseModel)
}

func NewUrlToDeepRouter(linkConverterService services.LinkConverterServiceInterface) UrlToDeepRouterInterface {
	return &UrlToDeepRouter{
		linkConverterService: linkConverterService,
	}
}
