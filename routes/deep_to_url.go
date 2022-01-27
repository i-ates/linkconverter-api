package routes

import (
	"encoding/json"
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
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(context.Request().Body).Decode(&jsonMap)

	if err != nil {
		return err
	}

	url := jsonMap["url"].(string) //bu burda mı olmalı bilemedim ilerde farklı bir field geldiğinde işimiz karışır

	response := deepToUrlRouter.linkConverterService.ConvertDeepToUrl(url)

	return context.JSON(http.StatusOK, response)
}

func NewDeepToUrlRouter(linkConverterService services.LinkConverterServiceInterface) DeepToUrlRouterInterface {
	return &DeepToUrlRouter{
		linkConverterService: linkConverterService,
	}
}
