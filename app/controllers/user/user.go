package user

import (
	"CWall/app/models"
	"CWall/app/services"
	"CWall/app/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type ChangeUserPasswordData struct {
	Account     string `json:"account"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ChangeUserPassword(c *gin.Context) {
	var data ChangeUserPasswordData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 用户是否存在
	user, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	//旧密码是否正确
	if data.OldPassword != user.Password {
		utils.JsonErrorResponse(c, 200507, "旧密码错误")
		return
	}
	//修改密码
	err = services.UpdateUserPassword(data.Account, data.NewPassword)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "修改失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type ChangeUserNameData struct {
	Account string `json:"account"`
	NewName string `json:"new_name"`
}

func ChangeUserName(c *gin.Context) {
	var data ChangeUserNameData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	//修改昵称
	err = services.UpdateUserName(data.Account, data.NewName)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "修改失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

// type ChangeUserAvatar struct {
// 	Account string `json:"account"`
// }

// func ChangeUserAvatar(c *gin.Context) {

// }
type BlockOtherData struct {
	Account      string `json:"account"`
	BlockAccount string `json:"block_account"`
}

func BlockOther(c *gin.Context) {
	var data BlockOtherData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 拉黑用户
	err = services.BlockUser(data.Account, data.BlockAccount)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "拉黑失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type SearchData struct {
	Account string `form:"account"`
	Keyword string `form:"keyword"`
}

func Search(c *gin.Context) {
	var data SearchData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	var filteredPost []models.Post
	postList, err := services.GetAllPostList()
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "表白列表获取失败")
		return
	}
	for _, post := range postList {
		if strings.Contains(post.PostContent, data.Keyword) {
			filteredPost = append(filteredPost, post)
		}
	}
	utils.JsonSuccess(c, filteredPost)
}

func GetHotRanking(c *gin.Context) {
	Account := c.Query("account")
	_, err := services.GetUserByAccount(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	loveList, err := services.GetAllLoveList()
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "获取数据失败")
		return
	}
	postIDCounts := make(map[int]int)
	for _, love := range loveList {
		postIDCounts[love.PostID]++
	}
	utils.JsonSuccess(c, postIDCounts)
}
