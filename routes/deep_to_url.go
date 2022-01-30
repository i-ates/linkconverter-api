package routes

import (
	"github.com/labstack/echo/v4"
	"linkconverter-api/models/requests"
	"linkconverter-api/services"
	"net/http"
)

type DeepToUrlRouterInterface interface {
	DeepToUrl(context echo.Context) error
}

type DeepToUrlRouter struct {
	linkConverterService services.LinkConverterServiceInterface
}

func (deepToUrlRouter *DeepToUrlRouter) DeepToUrl(context echo.Context) error {

	urlRequestModel := requests.UrlRequestModel{}

	err := context.Bind(&urlRequestModel)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	deepToUrlResponseModel, err := deepToUrlRouter.linkConverterService.ConvertDeepToUrl(urlRequestModel)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, deepToUrlResponseModel)
}

func NewDeepToUrlRouter(linkConverterService services.LinkConverterServiceInterface) DeepToUrlRouterInterface {
	return &DeepToUrlRouter{
		linkConverterService: linkConverterService,
	}
}
