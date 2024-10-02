package admin

import (
	"CWall/app/services"
	"CWall/app/utils"

	"github.com/gin-gonic/gin"
)

type GetAllReportListData struct {
	Account string `form:"account"`
	part    int    `formL:"part"`
}

func GetAllReportList(c *gin.Context) {
	var data GetAllReportListData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "用户不存在")
		return
	}
	if user.UserType != 1 {
		utils.JsonErrorResponse(c, 200501, "非管理员用户")
		return
	}
	report_list, err := services.GetAdminReviewList(data.part)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "获取失败")
		return
	}

	utils.JsonSuccess(c, report_list)
}

type ApprovalReportData struct {
	Account  string `json:"account"`
	PostID   int    `json:"post_id"`
	Approval int    `json:"approval"`
}

func ApprovalReport(c *gin.Context) {
	var data ApprovalReportData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "用户不存在")
		return
	}
	if user.UserType != 1 {
		utils.JsonErrorResponse(c, 200501, "非管理员用户")
		return
	}
	_, err = services.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	err = services.UpdateReportStatus(data.PostID, data.Approval)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "修改状态失败")
		return
	}
	if data.Approval == 1 {
		err = services.UpdateReportContent(data.PostID)
		if err != nil {
			utils.JsonErrorResponse(c, 200509, "修改举报内容失败")
			return
		}
		err = services.DeletePost(data.PostID)
		if err != nil {
			utils.JsonErrorResponse(c, 200509, "删除贴子失败")
			return
		}
	}
	utils.JsonSuccess(c, nil)
}
