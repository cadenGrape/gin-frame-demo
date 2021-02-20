package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func ApiResponse(c *gin.Context, code int, message string, data interface{})  {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}

func Index(c *gin.Context)  {
	ApiResponse(c, 0, "success", nil)
}
