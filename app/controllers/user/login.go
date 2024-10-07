package user

import (
	"CWall/app/models"
	"CWall/app/services"
	"CWall/app/utils"
	"github.com/gin-gonic/gin"
)


type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResData struct {
	Token string `json:"token"`
}

func Login_Post(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 500, "无效参数")
		return
	}
	var user *models.User
	user, err = services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200506, "用户不存在")
		return
	}
	if user.Status == 1 { //用户账号封禁
		utils.JsonErrorResponse(c, 200506, "封号中")
		return
	}

	flag := data.Password == user.Password
	if !flag {
		utils.JsonErrorResponse(c, 200507, "密码错误")
		return
	} else {
		token, err := utils.GenerateToken(data.Account)
		if err != nil {
			utils.JsonErrorResponse(c, 510, "生成令牌失败")
			return
		}
		res := LoginResData{Token: token}
		utils.JsonSuccess(c, res)
		return
	}
}
