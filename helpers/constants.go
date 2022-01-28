package helpers

const AppName = "LinkConverterApi"

var ProductDetailPageRegex = `https:\/\/www\.trendyol\.com\/[a-z0-9-]+\/[a-z0-9-]+-p-[0-9]+\?*[a-zA-Z]*=?[0-9]*&*[a-zA-Z]*=*[0-9]*`
var SearchPageRegex = `https:\\/\\/www\\.trendyol\\.com\\/sr\\?q=[a-zA-Z0-9%]+`
var OtherPagesRegex = `https:\\/\\/www\\.trendyol\\.com\\/`

var ProductDetailPageType = "ProductDetail"
var SearchPageType = "SearchPage"
var OtherPagesType = "OtherPages"
