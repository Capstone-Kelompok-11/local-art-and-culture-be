package response

import (
	"lokasani/helpers/errors"
	"net/http"

	"github.com/labstack/echo"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponse(c echo.Context, err error) error {
	return c.JSON(errors.GetCodeError(err), BaseResponse{
		Status:  false,
		Message: err.Error(),
		Data:    nil,
	})
}