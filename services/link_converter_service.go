package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"linkconverter-api/helpers"
	"linkconverter-api/models/responses"
	"linkconverter-api/parsers"
	"regexp"
)

type LinkConverterServiceInterface interface {
	ConvertDeepToUrl(context echo.Context) (responses.DeepToUrlResponseModel, error)
	ConvertUrlToDeep(context echo.Context) (responses.UrlToDeepResponseModel, error)
}

type LinkConverterService struct {
	urlParser parsers.UrlParserInterface
}

func (linkConverterService LinkConverterService) ConvertDeepToUrl(context echo.Context) (responses.DeepToUrlResponseModel, error) {
	var response responses.DeepToUrlResponseModel

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(context.Request().Body).Decode(&jsonMap)

	if err != nil {
		return response, err
	}

	url := jsonMap["deepLink"].(string)
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

	isUrlValid, pageType := linkConverterService.isUrlValidAndGetPageType(url)

	if isUrlValid == false {
		fmt.Println(pageType)
		return response, errors.New("url format is not correct")
	}

	deepLink := linkConverterService.convertUrl(url, pageType)
	if deepLink == "" {
		return response, errors.New("deepLink can not be empty")
	}

	response.DeepLink = deepLink
	return response, nil
}

func (linkConverterService LinkConverterService) isUrlValidAndGetPageType(url string) (bool, string) {
	r, _ := regexp.Compile(helpers.ProductDetailPageRegex)
	matched := r.MatchString(url)

	if matched == true {
		return matched, helpers.ProductDetailPageType
	}
	r, _ = regexp.Compile(helpers.SearchPageRegex)
	matched = r.MatchString(url)

	if matched {
		return matched, helpers.SearchPageType
	}

	r, _ = regexp.Compile(helpers.OtherPagesRegex)
	matched = r.MatchString(url)

	if matched {
		return matched, helpers.OtherPagesType
	}

	return false, ""

}

func (linkConverterService LinkConverterService) convertUrl(url string, pageType string) string {
	u := linkConverterService.urlParser.Parse(url, pageType)
	return u.BrandOrCategoryName

}

func NewLinkConverterService(urlParser parsers.UrlParserInterface) LinkConverterServiceInterface {
	return &LinkConverterService{
		urlParser: urlParser,
	}
}
