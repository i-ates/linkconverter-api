package services

import (
	_ "github.com/go-sql-driver/mysql"
	"linkconverter-api/builders"
	"linkconverter-api/models"
	"linkconverter-api/models/requests"
	"linkconverter-api/models/responses"
	"linkconverter-api/parsers"
)

type LinkConverterServiceInterface interface {
	ConvertDeepToUrl(deepLinkRequestModel requests.DeepLinkRequestModel) (responses.DeepToUrlResponseModel, error)
	ConvertUrlToDeep(urlRequestModel requests.UrlRequestModel) (responses.UrlToDeepResponseModel, error)
}

type LinkConverterService struct {
	urlParser  parsers.UrlParserInterface
	urlBuilder builders.UrlBuilderInterface
	dbBuilder  builders.DbBuilderInterface
}

func (linkConverterService LinkConverterService) ConvertDeepToUrl(deepLinkRequestModel requests.DeepLinkRequestModel) (responses.DeepToUrlResponseModel, error) {
	deepToUrlResponseModel := responses.DeepToUrlResponseModel{}

	parsedUrlModel, err := linkConverterService.urlParser.Parse(deepLinkRequestModel.DeepLink)
	linkConverterService.dbBuilder.DbConnection()
	if err != nil {
		linkConverterService.dbBuilder.InsertLogEvent(models.NewEvent(deepLinkRequestModel.DeepLink, ""))
		return deepToUrlResponseModel, err
	}

	linkConverterService.urlBuilder.BuildUrlUrl(&deepToUrlResponseModel, parsedUrlModel)

	linkConverterService.dbBuilder.InsertLogEvent(models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url))

	return deepToUrlResponseModel, nil
}

func (linkConverterService LinkConverterService) ConvertUrlToDeep(urlRequestModel requests.UrlRequestModel) (responses.UrlToDeepResponseModel, error) {
	urlToDeepResponseModel := responses.UrlToDeepResponseModel{}

	parsedUrlModel, err := linkConverterService.urlParser.Parse(urlRequestModel.Url)

	if err != nil {
		return urlToDeepResponseModel, err
	}

	linkConverterService.urlBuilder.BuildDeepUrl(&urlToDeepResponseModel, parsedUrlModel)

	return urlToDeepResponseModel, nil
}

func NewLinkConverterService(urlParser parsers.UrlParserInterface, urlBuilder builders.UrlBuilderInterface,
	dbBuilder builders.DbBuilderInterface) LinkConverterServiceInterface {
	return &LinkConverterService{
		urlParser:  urlParser,
		urlBuilder: urlBuilder,
		dbBuilder:  dbBuilder,
	}
}
