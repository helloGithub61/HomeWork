package user

import (
	"CWall/app/services"
	"CWall/app/utils"
	"time"
	"github.com/gin-gonic/gin"
)

func Reg_Get(c *gin.Context) {
	account := utils.RandomString(11)
	c.HTML(200, "reg.html", gin.H{
		"account": account,
	})
}

type RegDate struct {
	Account  string `json:"account" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	VerCode  string `json:"ver_code" binding:"required"`
}

func Reg_Post(c *gin.Context) {
	var data RegDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 参数校验
	//密码长度
	if len(data.Password) < 8 || len(data.Password) > 16 {
		utils.JsonErrorResponse(c, 200503, "密码长度必须在8-16位")
		return
	}
	//查看该邮箱是否注册过多账号
	num, err := services.GetUserNumByEmail(data.Email)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "获取数量失败")
		return
	}
	if num >= 2 {
		utils.JsonErrorResponse(c, 200501, "邮箱注册账号过多")
		return
	}
	// 验证码校验
	email, err := services.GetEmailByAddress(data.Email)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "获取验证码失败")
		return
	}
	// 验证码正确性校验
	if data.VerCode != email.Code {
		utils.JsonErrorResponse(c, 200504, "验证码错误")
		return
	}
	// 验证码时效校验
	nowTime := time.Now()
	duration := nowTime.Sub(email.Create_time)
	if duration > 5*time.Minute {
		utils.JsonErrorResponse(c, 500, "验证码过期")
		return
	}
	// 业务逻辑
	// 创建用户
	Key := utils.GenerateKey(data.Account, data.Email, data.Password)
	err = services.CreateUser(data.Account, data.Name, data.Email, data.Password, Key)
	if err != nil {
		utils.JsonErrorResponse(c, 200505, "注册失败")
		return
	}
	services.DeleteEmail(data.Email)
	utils.JsonSuccess(c, nil)
}
