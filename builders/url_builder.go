package builders

import (
	"linkconverter-api/helpers"
	"linkconverter-api/models"
	"linkconverter-api/models/responses"
	Url "net/url"
)

type UrlBuilderInterface interface {
	BuildDeepUrl(urlToDeepResponseModel *responses.UrlToDeepResponseModel, parsedUrlModel models.ParsedUrlModel)
	BuildUrlUrl(deepToUrlResponseModel *responses.DeepToUrlResponseModel, parsedUrlModel models.ParsedUrlModel)
}
type UrlBuilder struct {
}

func (urlBuilder UrlBuilder) BuildUrlUrl(deepToUrlResponseModel *responses.DeepToUrlResponseModel, parsedUrlModel models.ParsedUrlModel) {
	urlUrl := helpers.UrlBaseUrl

	if parsedUrlModel.PageType == helpers.ProductDetailPageType {
		urlUrl = urlUrl + helpers.UrlProductPageParamKey
	}
	if parsedUrlModel.PageType == helpers.SearchPagePageType {
		urlUrl = urlUrl + helpers.UrlSearchPageParamKey
	}
	if parsedUrlModel.PageType == helpers.OtherPagesPageType {
		deepToUrlResponseModel.Url = urlUrl
		return
	}
	anyParamAdded := false

	if parsedUrlModel.ContentId != "" {
		urlUrl = urlUrl + parsedUrlModel.ContentId + "?"
	}
	if parsedUrlModel.BoutiqueId != "" {
		urlUrl = urlUrl + helpers.UrlBoutiqueIdParamKey + "=" + parsedUrlModel.BoutiqueId
		anyParamAdded = true
	}
	if parsedUrlModel.MerchantId != "" {
		if anyParamAdded {
			urlUrl = urlUrl + "&" + helpers.UrlMerchantIdParamKey + "=" + parsedUrlModel.MerchantId
		} else {
			urlUrl = urlUrl + helpers.UrlMerchantIdParamKey + "=" + parsedUrlModel.MerchantId

		}
		anyParamAdded = true
	}
	if !anyParamAdded {
		urlUrl = urlUrl[:len(urlUrl)-1]
	}
	if parsedUrlModel.Q != "" {
		urlUrl = urlUrl + "?" + helpers.UrlProductQueryParamKey + "=" + Url.QueryEscape(parsedUrlModel.Q)
	}

	deepToUrlResponseModel.Url = urlUrl
}

func (urlBuilder UrlBuilder) BuildDeepUrl(urlToDeepResponseModel *responses.UrlToDeepResponseModel, parsedUrlModel models.ParsedUrlModel) {

	deepLinkUrl := helpers.DeepLinkBaseUrl

	if parsedUrlModel.PageType == helpers.OtherPagesPageType {
		deepLinkUrl = deepLinkUrl + helpers.DeepLinkOthersPageTypeParamKey
	} else {
		deepLinkUrl = deepLinkUrl + parsedUrlModel.PageType
	}

	if parsedUrlModel.ContentId != "" {
		deepLinkUrl = deepLinkUrl + "&" + helpers.DeepLinkContentIdParamKey + "=" + parsedUrlModel.ContentId
	}
	if parsedUrlModel.BoutiqueId != "" {
		deepLinkUrl = deepLinkUrl + "&" + helpers.DeepLinkCampaignIdParamKey + "=" + parsedUrlModel.BoutiqueId
	}
	if parsedUrlModel.MerchantId != "" {
		deepLinkUrl = deepLinkUrl + "&" + helpers.DeepLinkMerchantIdParamKey + "=" + parsedUrlModel.MerchantId
	}
	if parsedUrlModel.Q != "" {
		deepLinkUrl = deepLinkUrl + "&" + helpers.DeepLinkQueryParamKey + "=" + Url.QueryEscape(parsedUrlModel.Q)
	}

	urlToDeepResponseModel.DeepLink = deepLinkUrl

}

func NewUrlBuilder() UrlBuilderInterface {
	return &UrlBuilder{}
}
