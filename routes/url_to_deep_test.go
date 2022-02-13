package routes

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"linkconverter-api/mocks"
	"linkconverter-api/models/requests"
	"linkconverter-api/models/responses"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUrlToDeepRouter_UrlToDeep(t *testing.T) {
	e := echo.New()

	body := `{"url": "https://www.trendyol.com/casio/saat-p-1925865"}`
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	context := e.NewContext(req, recorder)

	urlToDeepResponseModel := responses.UrlToDeepResponseModel{DeepLink: "ty://?Page=Product&ContentId=1925865"}
	urlRequestModel := requests.UrlRequestModel{Url: "https://www.trendyol.com/casio/saat-p-1925865"}

	mockLinkConverterService := new(mocks.LinkConverterServiceInterface)
	mockLinkConverterService.On("ConvertUrlToDeep", urlRequestModel).Return(urlToDeepResponseModel, nil)

	urlToDeepRouter := NewUrlToDeepRouter(mockLinkConverterService)

	urlToDeepRouter.UrlToDeep(context)

	result := responses.UrlToDeepResponseModel{}
	json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Equal(t, urlToDeepResponseModel.DeepLink, result.DeepLink)
}

func TestNewUrlToDeepRouter(t *testing.T) {
	urlToDeepRouter := NewUrlToDeepRouter(nil)

	assert.NotNil(t, urlToDeepRouter)
}
