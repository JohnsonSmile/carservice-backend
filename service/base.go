package service

import (
	"carservice/service/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PARAM_ERROR,
			Msg:  response.MsgForCode(response.PARAM_ERROR),
		})
		return
	}
	c.JSON(http.StatusOK, &response.Response{
		Code: response.PARAM_ERROR,
		Msg:  errs.Error(),
	})
}
