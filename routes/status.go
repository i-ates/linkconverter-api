package routes

import (
	"github.com/labstack/echo/v4"
	"linkconverter-api/models/responses"
	"net/http"
)

type StatusRouterInterface interface {
	Status(context echo.Context) error
}

type StatusRouter struct {
}

// swagger:route GET /status Status getStatus
// responses:
//	200: StatusResponseModel
func (statusRouter *StatusRouter) Status(context echo.Context) error {
	responseModel := responses.StatusResponseModel{
		Status:      true,
		Description: "Always true",
	}

	return context.JSON(http.StatusOK, responseModel)
}

func NewStatusRouter() StatusRouterInterface {
	return &StatusRouter{}
}
