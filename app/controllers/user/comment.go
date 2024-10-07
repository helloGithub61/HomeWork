package user

import (
	"CWall/app/services"
	"CWall/app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateCommentData struct {
	Account string `json:"account"`
	Content string `json:"content"`
	Target  int    `json:"target"`
	PostID  int    `json:"post_id"`
}

func CreateComment(c *gin.Context) {
	var data CreateCommentData
	if err := c.ShouldBindJSON(&data); err != nil {

		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	// 1.用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.创建评论
	err = services.CreateComment(data.Account, data.Content, data.PostID, data.Target)
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "创建失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetCommentData struct {
	PostID int `form:"post_id"`
}

func GetComment(c *gin.Context) {
	var data GetCommentData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	fmt.Println(data.PostID)
	commentLst, err := services.GetCommentByPostID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "创建失败")
		return
	}
	utils.JsonSuccess(c, commentLst)
}
