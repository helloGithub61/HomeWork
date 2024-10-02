package controllers

import (
	"github.com/gin-gonic/gin"
)

func RootHandle(c *gin.Context) {
	c.Redirect(302, "/login")
}
