// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "linkconverter-api/models"
import responses "linkconverter-api/models/responses"

// UrlBuilderInterface is an autogenerated mock type for the UrlBuilderInterface type
type UrlBuilderInterface struct {
	mock.Mock
}

// BuildDeepUrl provides a mock function with given fields: urlToDeepResponseModel, parsedUrlModel
func (_m *UrlBuilderInterface) BuildDeepUrl(urlToDeepResponseModel *responses.UrlToDeepResponseModel, parsedUrlModel models.ParsedUrlModel) {
	_m.Called(urlToDeepResponseModel, parsedUrlModel)
}

// BuildUrlUrl provides a mock function with given fields: deepToUrlResponseModel, parsedUrlModel
func (_m *UrlBuilderInterface) BuildUrlUrl(deepToUrlResponseModel *responses.DeepToUrlResponseModel, parsedUrlModel models.ParsedUrlModel) {
	_m.Called(deepToUrlResponseModel, parsedUrlModel)
}
