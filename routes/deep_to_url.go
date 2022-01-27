package routes

import (
	"github.com/labstack/echo/v4"
	"linkconverter-api/services"
	"net/http"
)

type DeepToUrlRouterInterface interface {
	DeepToUrl(context echo.Context) error
}

type DeepToUrlRouter struct {
	linkConverterService services.LinkConverterServiceInterface
}

// swagger:route POST /deepToUrl DeepToUrl postUrlLink
// responses:
//	200: DeepToUrlResponseModel
func (deepToUrlRouter *DeepToUrlRouter) DeepToUrl(context echo.Context) error {

	response, err := deepToUrlRouter.linkConverterService.ConvertDeepToUrl(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, nil)
	}

	return context.JSON(http.StatusOK, response)
}

func NewDeepToUrlRouter(linkConverterService services.LinkConverterServiceInterface) DeepToUrlRouterInterface {
	return &DeepToUrlRouter{
		linkConverterService: linkConverterService,
	}
}
