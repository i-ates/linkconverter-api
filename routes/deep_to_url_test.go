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

func TestDeepToUrlRouter_DeepToUrl(t *testing.T) {
	e := echo.New()

	body := `{"deepLink": "ty://?Page=Product&ContentId=1925865"}`
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	context := e.NewContext(req, recorder)

	deepToUrlResponseModel := responses.DeepToUrlResponseModel{Url: "https://www.trendyol.com/brand/name-p-1925865"}
	deepLinkRequestModel := requests.DeepLinkRequestModel{DeepLink: "ty://?Page=Product&ContentId=1925865"}

	mockLinkConverterService := new(mocks.LinkConverterServiceInterface)
	mockLinkConverterService.On("ConvertDeepToUrl", deepLinkRequestModel).Return(deepToUrlResponseModel, nil)

	urlToDeepRouter := NewDeepToUrlRouter(mockLinkConverterService)

	urlToDeepRouter.DeepToUrl(context)

	result := responses.DeepToUrlResponseModel{}
	json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Equal(t, deepToUrlResponseModel.Url, result.Url)
}

func TestNewDeepToUrlRouter(t *testing.T) {
	deepToUrlRouter := NewDeepToUrlRouter(nil)

	assert.NotNil(t, deepToUrlRouter)
}
