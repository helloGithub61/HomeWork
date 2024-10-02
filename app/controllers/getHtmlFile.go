package controllers

import (
	"CWall/app/services"
	"CWall/app/utils"
	"github.com/gin-gonic/gin"
)

func SendHomePageHtml(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "My Webpage",
	})
}

func SendSpacePageHtml(c *gin.Context) {
	Account := c.Query("account")
	user, err := services.GetUserByAccount(Account)
	if err != nil {
		utils.JsonResponse(c, 200, 500, "获取用户信息失败", nil)
		return
	}
	if user.UserType == 0 {
		c.HTML(200, "space_user.html", gin.H{
			"title": "My Space",
			"user":  user, // 将user传递到模板中
		})
	} else if user.UserType == 1 {
		c.HTML(200, "space_admin.html", gin.H{
			"title": "My Space",
			"user":  user, // 将admin传递到模板中
		})
	} else {
		utils.JsonErrorResponse(c, 500, "用户类型错误")
	}
}
