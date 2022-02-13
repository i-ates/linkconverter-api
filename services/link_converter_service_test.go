package services

import (
	"github.com/stretchr/testify/assert"
	"linkconverter-api/builders"
	"linkconverter-api/mocks"
	"linkconverter-api/models"
	"linkconverter-api/models/requests"
	"linkconverter-api/models/responses"
	"linkconverter-api/parsers"
	"testing"
)

func TestLinkConverterService_ConvertUrlToDeep(t *testing.T) {
	t.Run("Get Page Response with only ContentId", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

	t.Run("Get Page Response with ContentId and MerchantId ", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865?&merchantId=105064",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&MerchantId=105064",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

	t.Run("Get Page Response with ContentId and BoutiqueId ", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865?boutiqueId=439892",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

	t.Run("Get Page Response with ContentId, BoutiqueId, MerchantId ", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865?boutiqueId=439892&merchantId=105064",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

	t.Run("Get Search Response with q=elbise", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/sr?q=elbise",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Search&Query=elbise",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

	t.Run("Get Search Response with q=%C3%BCt%C3%BC", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/sr?q=%C3%BCt%C3%BC",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Search&Query=%C3%BCt%C3%BC",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

	t.Run("Get Other Response", func(t *testing.T) {
		urlRequestModel := requests.UrlRequestModel{
			Url: "https://www.trendyol.com/Hesabim/#/Favoriler",
		}
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Home",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(urlRequestModel.Url, urlToDeepResponseModel.DeepLink)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertUrlToDeep(urlRequestModel)

		assert.Equal(t, urlToDeepResponseModel, result)
	})

}

func TestLinkConverterService_ConvertDeepToUrl(t *testing.T) {
	t.Run("Get Page Response with only ContentId", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

	t.Run("Get Page Response with ContentId and MerchantId", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&MerchantId=105064",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865?merchantId=105064",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

	t.Run("Get Page Response with ContentId and CampaignId", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

	t.Run("Get Page Response with ContentId, CampaignId, MerchantId", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

	t.Run("Get Search Response with q=elbise", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Search&Query=elbise",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/sr?q=elbise",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

	t.Run("Get Search Response with q=%C3%BCt%C3%BC", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Search&Query=%C3%BCt%C3%BC",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/sr?q=%C3%BCt%C3%BC",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

	t.Run("Get Other Response", func(t *testing.T) {
		deepLinkRequestModel := requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Favorites",
		}
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com",
		}

		urlParser := parsers.NewUrlParser()
		urlBuilder := builders.NewUrlBuilder()

		mockDbBuilder := new(mocks.DbBuilderInterface)
		dbEvent := models.NewEvent(deepLinkRequestModel.DeepLink, deepToUrlResponseModel.Url)
		mockDbBuilder.On("InsertLogEvent", dbEvent).Return(nil)

		linkConverterService := LinkConverterService{
			urlParser,
			urlBuilder,
			mockDbBuilder,
		}

		result, _ := linkConverterService.ConvertDeepToUrl(deepLinkRequestModel)

		assert.Equal(t, deepToUrlResponseModel, result)
	})

}

func TestNewLinkConverterService(t *testing.T) {
	linkConverterService := NewLinkConverterService(nil, nil, nil)

	assert.NotNil(t, linkConverterService)
}
