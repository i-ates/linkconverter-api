package parsers

import (
	"errors"
	"linkconverter-api/helpers"
	"linkconverter-api/models"
	"linkconverter-api/models/requests"
	Url "net/url"
	"regexp"
	"strings"
)

type UrlParserInterface interface {
	Parse(urlRequestModel requests.UrlRequestModel) (models.ParsedUrlModel, error)
}

type UrlParser struct {
}

func (urlParser *UrlParser) Parse(urlRequestModel requests.UrlRequestModel) (models.ParsedUrlModel, error) {
	parsedUrlModel := models.NewParsedUrlModel()

	parsedUrl, err := Url.Parse(urlRequestModel.Url)
	if err != nil {
		return parsedUrlModel, err
	}

	parsedUrlModel.UrlType, err = urlParser.GetUrlType(parsedUrl.Scheme)
	if err != nil {
		return parsedUrlModel, err
	}

	parsedUrlModel.PageType, err = urlParser.GetPageType(urlRequestModel.Url, parsedUrlModel.UrlType)
	if err != nil {
		return parsedUrlModel, err
	}

	urlParser.GetUrlUrlValues(&parsedUrlModel, parsedUrl)

	return parsedUrlModel, nil

}

func (urlParser *UrlParser) GetUrlType(scheme string) (string, error) {
	if strings.Contains(scheme, helpers.DeepLinkScheme) {
		return helpers.DeeplinkUrlType, nil
	}

	if strings.Contains(scheme, helpers.UrlScheme) {
		return helpers.UrlUrlType, nil
	}

	return helpers.NoneUrlType, errors.New("url type is not correct")
}

func (urlParser UrlParser) GetPageType(url string, urlType string) (string, error) {
	if urlType == helpers.UrlUrlType {
		return urlParser.GetPageTypeForUrl(url)
	}
	if urlType == helpers.DeeplinkUrlType {
		return urlParser.GetPageTypeForDeepLink(url)
	}

	return helpers.NonePageType, errors.New("url page type is invalid")

}

func (urlParser UrlParser) GetPageTypeForUrl(url string) (string, error) {
	r, _ := regexp.Compile(helpers.UrlProductDetailPageRegex)
	matched := r.MatchString(url)

	if matched {
		return helpers.ProductDetailPageType, nil
	}
	r, _ = regexp.Compile(helpers.UrlSearchPageRegex)
	matched = r.MatchString(url)

	if matched {
		return helpers.SearchPagePageType, nil
	}

	r, _ = regexp.Compile(helpers.UrlOtherPagesRegex)
	matched = r.MatchString(url)

	if matched {
		return helpers.OtherPagesPageType, nil
	}
	return helpers.NonePageType, errors.New("url pageType is invalid")

}

func (urlParser UrlParser) GetPageTypeForDeepLink(url string) (string, error) {
	r, _ := regexp.Compile(helpers.DeepLinkProductDetailPageRegex)
	matched := r.MatchString(url)

	if matched {
		return helpers.ProductDetailPageType, nil
	}
	r, _ = regexp.Compile(helpers.DeepLinkSearchPageRegex)
	matched = r.MatchString(url)

	if matched {
		return helpers.SearchPagePageType, nil
	}

	r, _ = regexp.Compile(helpers.DeepLinkOtherPagesRegex)
	matched = r.MatchString(url)

	if matched {
		return helpers.OtherPagesPageType, nil
	}
	return helpers.NonePageType, errors.New("url pageType is not valid")

}

func (urlParser UrlParser) GetUrlValues(parsedUrlModel models.ParsedUrlModel, parsedUrl *Url.URL) {
	if parsedUrlModel.UrlType == helpers.UrlUrlType {
		urlParser.GetUrlUrlValues(&parsedUrlModel, parsedUrl)
	}

	if parsedUrlModel.UrlType == helpers.DeeplinkUrlType {
		urlParser.GetDeepLinkValues(&parsedUrlModel, parsedUrl)
	}
}

func (urlParser UrlParser) GetUrlUrlValues(parsedUrlModel *models.ParsedUrlModel, parsedUrl *Url.URL) {
	if parsedUrlModel.PageType == helpers.ProductDetailPageType {
		sPath := strings.Split(parsedUrl.Path, "-")
		parsedUrlModel.ContentId = sPath[len(sPath)-1]

		if boutiqueId, ok := parsedUrl.Query()[helpers.UrlBoutiqueIdParamKey]; ok {
			parsedUrlModel.BoutiqueId = strings.Join(boutiqueId, "")
		}

		if merchantId, ok := parsedUrl.Query()[helpers.UrlMerchantIdParamKey]; ok {
			parsedUrlModel.MerchantId = strings.Join(merchantId, "")
		}
	}

	if parsedUrlModel.PageType == helpers.SearchPagePageType {
		if q, ok := parsedUrl.Query()[helpers.UrlProductQueryParamKey]; ok {
			parsedUrlModel.Q = strings.Join(q, "")
		}
	}
}

func (urlParser *UrlParser) GetDeepLinkValues(parsedUrlModel *models.ParsedUrlModel, parsedUrl *Url.URL) {
	if parsedUrlModel.PageType == helpers.SearchPagePageType {
		if q, ok := parsedUrl.Query()[helpers.DeepLinkQueryParamKey]; ok {
			parsedUrlModel.Q = strings.Join(q, "")
		}
	}

	if parsedUrlModel.PageType == helpers.ProductDetailPageType {
		if q, ok := parsedUrl.Query()[helpers.DeepLinkContentIdParamKey]; ok {
			parsedUrlModel.ContentId = strings.Join(q, "")
		}
		if q, ok := parsedUrl.Query()[helpers.DeepLinkCampaignIdParamKey]; ok {
			parsedUrlModel.BoutiqueId = strings.Join(q, "")
		}
		if q, ok := parsedUrl.Query()[helpers.DeepLinkMerchantIdParamKey]; ok {
			parsedUrlModel.MerchantId = strings.Join(q, "")
		}
	}
}

func NewUrlParser() UrlParserInterface {
	return &UrlParser{}
}
