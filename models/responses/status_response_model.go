package responses

// swagger:model StatusResponseModel
type StatusResponseModel struct {
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
