package routes

import (
	"encoding/json"
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

// swagger:route POST /urlToDeep UrlToDeep postUrlLink
// responses:
//	200: UrlToDeepResponseModel
func (urlToDeepRouter *UrlToDeepRouter) UrlToDeep(context echo.Context) error {
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(context.Request().Body).Decode(&jsonMap)

	if err != nil {
		return err
	}

	url := jsonMap["url"].(string) //bu burda mı olmalı bilemedim ilerde farklı bir field geldiğinde işimiz karışır

	response := urlToDeepRouter.linkConverterService.ConvertDeepToUrl(url)

	return context.JSON(http.StatusOK, response)
}

func NewUrlToDeepRouter(linkConverterService services.LinkConverterServiceInterface) UrlToDeepRouterInterface {
	return &UrlToDeepRouter{
		linkConverterService: linkConverterService,
	}
}
