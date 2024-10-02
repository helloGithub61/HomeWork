package admin

import (
	"CWall/app/services"
	"CWall/app/utils"

	"github.com/gin-gonic/gin"
)

type GetAllUserInfoData struct {
	Account string `json:"account"`
}
type UserInfoResData struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

func GetAllUserInfo(c *gin.Context) {
	var data GetAllUserInfoData
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
	if user.UserType != 1 {
		utils.JsonErrorResponse(c, 200507, "非管理员身份")
		return
	}
	userList, err := services.GetAllUserList()
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "获取用户信息失败")
		return
	}
	var filteredItem []UserInfoResData
	for _, user := range userList {
		if user.UserType == 0 {
			item := UserInfoResData{
				Account:  user.Account,
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
				Status:   user.Status,
			}
			filteredItem = append(filteredItem, item)
		}
	}
	utils.JsonSuccess(c, filteredItem)

}

type LockUserData struct {
	Account    string `json:"account"`
	BanAccount string `json:"ban_account"`
}

func LockUser(c *gin.Context) {
	var data LockUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	if user.UserType != 1 {
		utils.JsonErrorResponse(c, 200507, "非管理员身份")
		return
	}
	err = services.UpdateUserStatus(data.BanAccount, 1)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "封禁失败")
		return
	}
	utils.JsonSuccess(c, nil)
}
func UnlockUser(c *gin.Context) {
	var data LockUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	if user.UserType != 1 {
		utils.JsonErrorResponse(c, 200507, "非管理员身份")
		return
	}
	err = services.UpdateUserStatus(data.BanAccount, 1)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "解封失败")
		return
	}
	utils.JsonSuccess(c, nil)
}
