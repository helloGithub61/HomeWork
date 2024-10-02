package utils

import (
	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
func JsonSuccess(c *gin.Context, data interface{}) {
	JsonResponse(c, 200, 200, "success", data)
}

func JsonErrorResponse(c *gin.Context, code int, msg string) {
	JsonResponse(c, 200, code, msg, nil)
}
