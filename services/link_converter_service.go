package services

import "linkconverter-api/models/responses"

type LinkConverterServiceInterface interface {
	ConvertDeepToUrl(url string) responses.DeepToUrlResponseModel
	ConvertUrlToDeep(url string) responses.UrlToDeepResponseModel
}

type LinkConverterService struct {
}

func (linkConverterService LinkConverterService) ConvertDeepToUrl(url string) responses.DeepToUrlResponseModel {
	var response responses.DeepToUrlResponseModel

	response.Url = url
	return response
}

func (linkConverterService LinkConverterService) ConvertUrlToDeep(url string) responses.UrlToDeepResponseModel {
	var response responses.UrlToDeepResponseModel

	response.DeepLink = url
	return response
}

func NewLinkConverterService() LinkConverterServiceInterface {
	return &LinkConverterService{}
}
