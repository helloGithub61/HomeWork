package user

import (
	"CWall/app/models"
	"CWall/app/services"
	"CWall/app/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type CreatePostData struct {
	Account  string `json:"account"`
	Content  string `json:"content"`
	ShowType int    `json:"show_type"`
}

func CreatePost(c *gin.Context) {
	var data CreatePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	// 业务逻辑
	// 1.用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.创建帖子
	err = services.CreatePost(data.Account, data.Content, data.ShowType)
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "创建失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetPostListResponse struct {
	PostList []models.Post `json:"post_list"`
}
type GetPostListData struct {
	Part int `json:"part"`
}

func GetPostList(c *gin.Context) {
	var data GetPostListData
	c.ShouldBindJSON(&data)
	var postList []models.Post
	postList, err := services.GetHomePostList(data.Part)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"post_list": postList,
	})
}

type GetPersonPostListData struct {
	Account string `json:"account"`
}

func GetPersonPostList(c *gin.Context) {
	var data GetPersonPostListData
	c.ShouldBindJSON(&data)
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	var postList []models.Post
	postList, err := services.GetPostListByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"post_list": postList,
	})
}

type UpdatePostData struct {
	Account string `json:"account"`
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var data UpdatePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	// 业务逻辑
	// 1.用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.帖子是否存在
	post, err := services.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	// 3.帖子是否属于用户
	if post.Account != data.Account {
		utils.JsonErrorResponse(c, 200510, "无权限")
		return
	}
	// 4.更新帖子
	err = services.UpdatePostContent(data.PostID, data.Content)
	if err != nil {
		utils.JsonErrorResponse(c, 200510, "更新失败")
		return
	}
	utils.JsonSuccess(c, nil)

}

type DeletePostData struct {
	Account string `form:"account"`
	PostID  int    `form:"post_id"`
}

func DeletePost(c *gin.Context) {
	var data DeletePostData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	// 业务逻辑
	// 1.用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.帖子是否存在
	post, err := services.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	// 3.帖子是否属于用户
	if data.Account != post.Account {
		utils.JsonErrorResponse(c, 200510, "无权限")
		return
	}
	// 4.删除帖子和更新举报记录
	err = services.DeletePost(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "删除失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type ReportPostData struct {
	Account string `json:"account"`
	PostID  int    `json:"post_id"`
	Reason  string `json:"reason"`
}

func ReportPost(c *gin.Context) {
	var data ReportPostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	// 业务逻辑
	// 1.用户是否存在
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.帖子是否存在
	_, err = services.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	// 3.举报帖子
	err = services.ReportPost(data.Account, data.PostID, data.Reason)
	if err != nil {
		utils.JsonErrorResponse(c, 200512, "举报失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetReportListResponse struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
}

func GetReportList(c *gin.Context) {
	Account := utils.GetAccountByToken(c.GetHeader("Authorization"))
	_, err := services.GetUserByAccount(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 业务逻辑
	// 1.获取举报列表
	var reportList []models.Report
	reportList, err = services.GetReportList(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200513, "获取失败")
		return
	}
	var reportListResponse []GetReportListResponse
	for _, report := range reportList {
		// 2.获取帖子内容
		post, err := services.GetPostByID(report.PostID)
		if err != nil {
			utils.JsonErrorResponse(c, 200514, "获取失败")
			return
		}
		// 3.返回帖子内容
		reportListResponse = append(reportListResponse, GetReportListResponse{
			PostID:  post.ID,
			Content: post.PostContent,
			Reason:  report.Reason,
			Status:  report.Status,
		})
	}

	utils.JsonSuccess(c, gin.H{
		"report_list": reportListResponse,
	})
}

type LovePostData struct {
	Account string `json:"account"`
	PostID  int    `json:"post_id"`
}

func LikePost(c *gin.Context) {
	var data LovePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	data.Account = utils.GetAccountByToken(c.GetHeader("Authorization"))
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	_, err = services.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	love, err := services.GetLoveByAccAndPID(data.Account, data.PostID)
	if err == nil {
		love.Like = 1 - love.Like
		res := services.UpdateLike(*love)
		if res != nil {
			utils.JsonErrorResponse(c, 200500, "更新点赞失败")
		}
	} else {
		services.CreateLove(models.Love{
			Account: data.Account,
			PostID:  data.PostID,
			Like:    1,
		})
		utils.JsonErrorResponse(c, 200500, "创建点赞失败")
	}
}

func CollectPost(c *gin.Context) {
	var data LovePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	_, err := services.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	_, err = services.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	love, err := services.GetLoveByAccAndPID(data.Account, data.PostID)
	if err == nil {
		love.Collect = 1 - love.Collect
		res := services.UpdateCollect(*love)
		if res != nil {
			utils.JsonErrorResponse(c, 200500, "更新收藏失败")
		}
	} else {
		services.CreateLove(models.Love{
			Account: data.Account,
			PostID:  data.PostID,
			Collect: 1,
		})
		utils.JsonErrorResponse(c, 200500, "创建收藏失败")
	}
}

func GetLikePostList(c *gin.Context) {
	Account := utils.GetAccountByToken(c.GetHeader("Authorization"))
	_, err := services.GetUserByAccount(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	loveList, err := services.GetLoveList(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "获取点赞失败")
		return
	}
	var resList []models.Post
	for _, love := range loveList {
		if love.Like == 1 {
			post, err := services.GetPostByID(love.PostID)
			if err == nil {
				resList = append(resList, *post)
			}
		}
	}
	utils.JsonSuccess(c, resList)
}
func GetCollectPostList(c *gin.Context) {
	Account := utils.GetAccountByToken(c.GetHeader("Authorization"))
	_, err := services.GetUserByAccount(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	loveList, err := services.GetLoveList(Account)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "获取收藏失败")
		return
	}
	var resList []models.Post
	for _, love := range loveList {
		if love.Collect == 1 {
			post, err := services.GetPostByID(love.PostID)
			if err == nil {
				resList = append(resList, *post)
			}
		}
	}
	utils.JsonSuccess(c, resList)
}
func ChangeUserAvatar(c *gin.Context) {
	Account := utils.GetAccountByToken(c.GetHeader("Authorization"))
	file, _ := c.FormFile("file")
	timestamp := time.Now().Unix()
	dst := "./asset/pic/" + Account + "_" + fmt.Sprintf("%d", timestamp) + ".png"
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "头像上传失败")
		return
	}
	err = services.UpdateUserImage(Account, dst)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "头像更换失败")
		return
	}
	utils.JsonSuccess(c, nil)
}
