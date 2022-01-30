package parsers

import (
	"linkconverter-api/models"
	"log"
	Url "net/url"
	"strconv"
	"strings"
)

type UrlParserInterface interface {
	Parse(url string, pageType string) models.ParsedUrlModel
}

type UrlParser struct {
}

func (urlParser *UrlParser) Parse(url string, pageType string) models.ParsedUrlModel {
	model := models.NewParsedUrlModel()

	parsedUrl, _ := Url.Parse(url)

	if pageType == "ProductDetail" {
		urlParser.ParseProductDetailPage(*parsedUrl, &model)
	}

	return model

}

func (urlParser *UrlParser) ParseProductDetailPage(parsedUrl Url.URL, parsedUrlModel *models.ParsedUrlModel) {
	pathSplinted := strings.Split(parsedUrl.Path, "/")
	productNameAndContentId := strings.Split(pathSplinted[2], "-p-")

	brandOrCategoryName := pathSplinted[1]
	productName := productNameAndContentId[0]
	contentId := productNameAndContentId[1]

	parsedUrlModel.BrandOrCategoryName = brandOrCategoryName
	parsedUrlModel.ProductName = productName
	parsedUrlModel.ContentId, _ = strconv.Atoi(contentId)

	if parsedUrl.RawQuery != "" {
		rawQuerySplinted := strings.Split(parsedUrl.RawQuery, "&")
		if len(rawQuerySplinted) > 0 {
			rawQueryMap := make(map[string]int)
			for i := 0; i < len(rawQuerySplinted); i++ {
				rawQueryElementSplinted := strings.Split(rawQuerySplinted[i], "=")
				rawQueryMap[rawQueryElementSplinted[0]], _ = strconv.Atoi(rawQueryElementSplinted[1])
			}
			//mapledikten sonra if elssiz model ile eşleştirme yapılır mı bilemedim

		}
	}
	log.Fatal("x")

}

func NewUrlParser() UrlParserInterface {
	return &UrlParser{}
}
