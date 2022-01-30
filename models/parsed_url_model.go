package models

type ParsedUrlModel struct {
	PageType   string
	UrlType    string
	Q          string
	ContentId  string
	BoutiqueId string
	MerchantId string
}

func NewParsedUrlModel() ParsedUrlModel {
	return ParsedUrlModel{}
}
