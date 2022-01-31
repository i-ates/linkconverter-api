package builders

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"linkconverter-api/helpers"
	"linkconverter-api/models"
	"linkconverter-api/models/responses"
	"testing"
)

var deepLinkCases = []struct {
	parsedUrlModel                 models.ParsedUrlModel
	expectedUrlToDeepResponseModel responses.UrlToDeepResponseModel
}{
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.UrlUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
			MerchantId: "105064",
		},
		expectedUrlToDeepResponseModel: responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:  helpers.ProductDetailPageType,
			UrlType:   helpers.UrlUrlType,
			ContentId: "1925865",
		},
		expectedUrlToDeepResponseModel: responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.UrlUrlType,
			ContentId:  "1925865",
			MerchantId: "105064",
		},
		expectedUrlToDeepResponseModel: responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&MerchantId=105064",
		},
	},

	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.UrlUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
		},
		expectedUrlToDeepResponseModel: responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.SearchPagePageType,
			UrlType:  helpers.UrlUrlType,
			Q:        "elbise",
		},
		expectedUrlToDeepResponseModel: responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Search&Query=elbise",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.OtherPagesPageType,
			UrlType:  helpers.UrlUrlType,
		},
		expectedUrlToDeepResponseModel: responses.UrlToDeepResponseModel{
			DeepLink: "ty://?Page=Home",
		},
	},
}

var urlCases = []struct {
	parsedUrlModel                 models.ParsedUrlModel
	expectedDeepToUrlResponseModel responses.DeepToUrlResponseModel
}{
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.DeeplinkUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
			MerchantId: "105064",
		},
		expectedDeepToUrlResponseModel: responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:  helpers.ProductDetailPageType,
			UrlType:   helpers.DeeplinkUrlType,
			ContentId: "1925865",
		},
		expectedDeepToUrlResponseModel: responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.DeeplinkUrlType,
			ContentId:  "1925865",
			MerchantId: "105064",
		},
		expectedDeepToUrlResponseModel: responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865?merchantId=105064",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.DeeplinkUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
		},
		expectedDeepToUrlResponseModel: responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.SearchPagePageType,
			UrlType:  helpers.DeeplinkUrlType,
			Q:        "elbise",
		},
		expectedDeepToUrlResponseModel: responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com/sr?q=elbise",
		},
	},
	{
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.OtherPagesPageType,
			UrlType:  helpers.DeeplinkUrlType,
		},
		expectedDeepToUrlResponseModel: responses.DeepToUrlResponseModel{
			Url: "https://www.trendyol.com",
		},
	},
}

var urlBuilder = UrlBuilder{}

func TestShouldBuildDeepLinks(t *testing.T) {
	for _, c := range deepLinkCases {
		urlToDeepResponseModel := responses.UrlToDeepResponseModel{}
		t.Run(fmt.Sprintf("%s", c.expectedUrlToDeepResponseModel.DeepLink), func(t *testing.T) {
			urlBuilder.BuildDeepUrl(&urlToDeepResponseModel, c.parsedUrlModel)
			assert.Equal(t, c.expectedUrlToDeepResponseModel, urlToDeepResponseModel)
		})
	}
}

func TestShouldBuildUrlLinks(t *testing.T) {
	for _, c := range urlCases {
		deepToUrlResponseModel := responses.DeepToUrlResponseModel{}
		t.Run(fmt.Sprintf("%s", c.expectedDeepToUrlResponseModel.Url), func(t *testing.T) {
			urlBuilder.BuildUrlUrl(&deepToUrlResponseModel, c.parsedUrlModel)
			assert.Equal(t, c.expectedDeepToUrlResponseModel, deepToUrlResponseModel)
		})
	}
}

func TestNewUrlBuilder(t *testing.T) {
	t.Run("GetUrlBuilder", func(t *testing.T) {
		config := NewUrlBuilder()

		assert.NotZero(t, config)
	})
}
