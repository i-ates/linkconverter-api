package parsers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"linkconverter-api/helpers"
	"linkconverter-api/models"
	"linkconverter-api/models/requests"
	"testing"
)

var urlCases = []struct {
	urlRequestModel requests.UrlRequestModel
	parsedUrlModel  models.ParsedUrlModel
	err             error
}{
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865?boutiqueId=439892&merchantId=105064",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.UrlUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
			MerchantId: "105064",
		},
		err: nil,
	},
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:  helpers.ProductDetailPageType,
			UrlType:   helpers.UrlUrlType,
			ContentId: "1925865",
		},
		err: nil,
	},
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865?&merchantId=105064",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.UrlUrlType,
			ContentId:  "1925865",
			MerchantId: "105064",
		},
		err: nil,
	},
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/casio/saat-p-1925865?boutiqueId=439892",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.UrlUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
		},
		err: nil,
	},
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/sr?q=elbise",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.SearchPagePageType,
			UrlType:  helpers.UrlUrlType,
			Q:        "elbise",
		},
		err: nil,
	},
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/Hesabim/#/Favoriler",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.OtherPagesPageType,
			UrlType:  helpers.UrlUrlType,
		},
		err: nil,
	},
	{
		urlRequestModel: requests.UrlRequestModel{
			Url: "https://www.trendyol.com/Hesabim/#/Siparislerim",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.OtherPagesPageType,
			UrlType:  helpers.UrlUrlType,
		},
		err: nil,
	},
}

var deepLinkCases = []struct {
	deepLinkRequestModel requests.DeepLinkRequestModel
	parsedUrlModel       models.ParsedUrlModel
	err                  error
}{
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.DeeplinkUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
			MerchantId: "105064",
		},
		err: nil,
	},
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:  helpers.ProductDetailPageType,
			UrlType:   helpers.DeeplinkUrlType,
			ContentId: "1925865",
		},
		err: nil,
	},
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&MerchantId=105064",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.DeeplinkUrlType,
			ContentId:  "1925865",
			MerchantId: "105064",
		},
		err: nil,
	},
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Product&ContentId=1925865&CampaignId=439892",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType:   helpers.ProductDetailPageType,
			UrlType:    helpers.DeeplinkUrlType,
			ContentId:  "1925865",
			BoutiqueId: "439892",
		},
		err: nil,
	},
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Search&Query=elbise",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.SearchPagePageType,
			UrlType:  helpers.DeeplinkUrlType,
			Q:        "elbise",
		},
		err: nil,
	},
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Favorites",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.OtherPagesPageType,
			UrlType:  helpers.DeeplinkUrlType,
		},
		err: nil,
	},
	{
		deepLinkRequestModel: requests.DeepLinkRequestModel{
			DeepLink: "ty://?Page=Orders",
		},
		parsedUrlModel: models.ParsedUrlModel{
			PageType: helpers.OtherPagesPageType,
			UrlType:  helpers.DeeplinkUrlType,
		},
		err: nil,
	},
}

var urlParser = UrlParser{}

func TestShouldParseUrls(t *testing.T) {
	for _, c := range urlCases {
		t.Run(fmt.Sprintf("%s", c.urlRequestModel), func(t *testing.T) {
			result, err := urlParser.Parse(c.urlRequestModel.Url)
			assert.Equal(t, c.parsedUrlModel, result)
			assert.Equal(t, err, c.err)
		})
	}
}

func TestShouldParseDeepLinks(t *testing.T) {
	for _, c := range deepLinkCases {
		t.Run(fmt.Sprintf("%s", c.deepLinkRequestModel), func(t *testing.T) {
			result, err := urlParser.Parse(c.deepLinkRequestModel.DeepLink)
			assert.Equal(t, c.parsedUrlModel, result)
			assert.Equal(t, err, c.err)
		})
	}
}

func TestNewUrlParser(t *testing.T) {
	t.Run("GetUrlParser", func(t *testing.T) {
		config := NewUrlParser()

		assert.NotZero(t, config)
	})
}
