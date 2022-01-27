package services

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"linkconverter-api/models/responses"
)

type LinkConverterServiceInterface interface {
	ConvertDeepToUrl(context echo.Context) (responses.DeepToUrlResponseModel, error)
	ConvertUrlToDeep(context echo.Context) (responses.UrlToDeepResponseModel, error)
}

type LinkConverterService struct {
}

func (linkConverterService LinkConverterService) ConvertDeepToUrl(context echo.Context) (responses.DeepToUrlResponseModel, error) {
	var response responses.DeepToUrlResponseModel

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(context.Request().Body).Decode(&jsonMap)

	if err != nil {
		return response, err
	}

	url := jsonMap["url"].(string)
	if url == "" {
		return response, errors.New("deepLink can not be empty")
	}

	response.Url = url

	return response, nil
}

func (linkConverterService LinkConverterService) ConvertUrlToDeep(context echo.Context) (responses.UrlToDeepResponseModel, error) {
	var response responses.UrlToDeepResponseModel

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(context.Request().Body).Decode(&jsonMap)

	if err != nil {
		return response, err
	}

	url := jsonMap["url"].(string)
	if url == "" {
		return response, errors.New("url can not be empty")
	}

	response.DeepLink = url
	return response, nil
}

func NewLinkConverterService() LinkConverterServiceInterface {
	return &LinkConverterService{}
}
