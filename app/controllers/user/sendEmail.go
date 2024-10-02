package user

import (
	"CWall/app/services"
	"CWall/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailData struct {
	Email string `json:"email"`
}

func SendEmailCode(c *gin.Context) {
	var data EmailData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 500, "error1")
		return
	}
	verify := utils.RandomString(5)
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "表白墙 <2109702578@qq.com>"
	// 设置接收方的邮箱
	e.To = []string{data.Email}
	//设置主题
	e.Subject = "测试"
	//设置文件发送的内容
	e.Text = []byte("欢迎使用表白墙，你的验证码是" + verify)
	//设置服务器相关的配置
	e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2109702578@qq.com", "elqkpienhyirbiaf", "smtp.qq.com"))
	_, err := services.GetEmailByAddress(data.Email)
	if err == nil {
		services.UpdateEmail(data.Email, verify)
	} else {
		err = services.CreateEmail(data.Email, verify)
	}

	if err != nil {
		utils.JsonErrorResponse(c, 500, "error2")
	}
	utils.JsonSuccess(c, nil)
}
