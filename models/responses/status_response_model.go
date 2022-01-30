package responses

type StatusResponseModel struct {
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
