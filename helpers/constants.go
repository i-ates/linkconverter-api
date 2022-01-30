package helpers

const UrlProductDetailPageRegex = `https:\/\/www\.trendyol\.com\/[a-z0-9-]+\/[a-z0-9-]+-p-[0-9]+`
const UrlSearchPageRegex = `https:\/\/www\.trendyol\.com\/sr\?q=[a-zA-Z0-9%]+`
const UrlOtherPagesRegex = `https:\/\/www\.trendyol\.com\/`

const DeepLinkProductDetailPageRegex = `ty:\/\/\?Page=[a-zA-Z]+&ContentId=[0-9]+`
const DeepLinkSearchPageRegex = `ty:\/\/\?Page=Search&Query=[a-zA-Z0-9%]+`
const DeepLinkOtherPagesRegex = `ty:\/\/\?Page=[a-zA-Z]+`

const ProductDetailPageType = "Product"
const SearchPagePageType = "Search"
const OtherPagesPageType = "Other"
const NonePageType = "None"

const DeepLinkScheme = "ty"
const DeepLinkBaseUrl = DeepLinkScheme + "://?Page="
const UrlScheme = "http"
const UrlBaseUrl = "https://www.trendyol.com"

const DeeplinkUrlType = "Deeplink"
const UrlUrlType = "Url"
const NoneUrlType = "None"

const UrlBoutiqueIdParamKey = "boutiqueId"
const UrlMerchantIdParamKey = "merchantId"
const UrlProductQueryParamKey = "q"

const DeepLinkContentIdParamKey = "ContentId"
const DeepLinkCampaignIdParamKey = "CampaignId"
const DeepLinkMerchantIdParamKey = "MerchantId"
const DeepLinkQueryParamKey = "Query"
const DeepLinkOthersPageTypeParamKey = "Home"

const UrlProductPageParamKey = "/brand/name-p-"
const UrlSearchPageParamKey = "/sr?"
