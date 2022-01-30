package models

// swagger:model ParsedUrlModel
type ParsedUrlModel struct {
	BrandOrCategoryName string
	ProductName         string
	ContentId           int
	BoutiqueId          int
	merchantId          int
}

func NewParsedUrlModel() ParsedUrlModel {
	return ParsedUrlModel{}
}
