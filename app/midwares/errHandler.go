package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code": 404,
		"msg":  "not found",
		"data": nil,
	})
}
