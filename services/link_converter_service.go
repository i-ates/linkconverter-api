package services

import (
	"linkconverter-api/builders"
	"linkconverter-api/models/requests"
	"linkconverter-api/models/responses"
	"linkconverter-api/parsers"
)

type LinkConverterServiceInterface interface {
	ConvertDeepToUrl(urlRequestModel requests.UrlRequestModel) (responses.DeepToUrlResponseModel, error)
	ConvertUrlToDeep(urlRequestModel requests.UrlRequestModel) (responses.UrlToDeepResponseModel, error)
}

type LinkConverterService struct {
	urlParser  parsers.UrlParserInterface
	urlBuilder builders.UrlBuilderInterface
}

func (linkConverterService LinkConverterService) ConvertDeepToUrl(urlRequestModel requests.UrlRequestModel) (responses.DeepToUrlResponseModel, error) {
	deepToUrlResponseModel := responses.DeepToUrlResponseModel{}

	parsedUrlModel, err := linkConverterService.urlParser.Parse(urlRequestModel)

	if err != nil {
		return deepToUrlResponseModel, err
	}

	linkConverterService.urlBuilder.BuildUrlUrl(&deepToUrlResponseModel, parsedUrlModel)

	return deepToUrlResponseModel, nil
}

func (linkConverterService LinkConverterService) ConvertUrlToDeep(urlRequestModel requests.UrlRequestModel) (responses.UrlToDeepResponseModel, error) {
	urlToDeepResponseModel := responses.UrlToDeepResponseModel{}

	parsedUrlModel, err := linkConverterService.urlParser.Parse(urlRequestModel)

	if err != nil {
		return urlToDeepResponseModel, err
	}

	linkConverterService.urlBuilder.BuildDeepUrl(&urlToDeepResponseModel, parsedUrlModel)

	return urlToDeepResponseModel, nil
}

func NewLinkConverterService(urlParser parsers.UrlParserInterface, urlBuilder builders.UrlBuilderInterface) LinkConverterServiceInterface {
	return &LinkConverterService{
		urlParser:  urlParser,
		urlBuilder: urlBuilder,
	}
}
