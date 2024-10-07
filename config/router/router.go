package router

import (
	"CWall/app/controllers"
	"CWall/app/controllers/admin"
	"CWall/app/controllers/user"
	"CWall/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	per := api.Group("/user").Use(middleware.TokenMiddleware)
	ad := api.Group("/admin").Use(middleware.TokenMiddleware)
	// 设置路由
	{
		r.GET("/", controllers.RootHandle)                   //根目录，后续跳转
		api.POST("/login", user.Login_Post)                  //登录操作
		api.POST("/verify", user.SendEmailCode)              //获取验证码
		api.POST("/reg", user.Reg_Post)                      //注册操作
		per.GET("/search", user.Search)                      //搜索表白
		per.POST("/post", user.CreatePost)                   //发送表白
		per.POST("/next-page", user.GetPostList)             //获取主页面表白
		per.POST("/per-post", user.GetPersonPostList)        //获取自己发送的表白
		per.DELETE("/post", user.DeletePost)                 //删除表白
		per.POST("/report-post", user.ReportPost)            //举报表白
		per.GET("/report-post", user.GetReportList)          //获取举报
		per.PUT("/post", user.UpdatePost)                    //修改表白
		per.POST("/like", user.LikePost)                     //点赞表白
		per.POST("/collect", user.CollectPost)               //收藏表白
		per.GET("/like", user.GetLikePostList)               //获取点赞
		per.GET("/collect", user.GetCollectPostList)         //获取收藏
		per.GET("/hot-rank", user.GetHotRanking)             //获取热度排行
		per.GET("/revise-avatar", user.ChangeUserAvatar)     //修改头像
		per.PUT("/revise-name", user.ChangeUserName)         //修改昵称
		per.PUT("/revise-password", user.ChangeUserPassword) //修改密码
		per.POST("/comment", user.CreateComment)             //发送评论
		per.GET("/comment", user.GetComment)                 //获取评论
		per.POST("/block", user.BlockOther)                  //拉黑
		ad.GET("/report", admin.GetAllReportList)            //管理员获取举报列表
		ad.POST("/report", admin.ApprovalReport)             //管理员审核举报
		ad.POST("/userInfo", admin.GetAllUserInfo)           //管理员查看用户信息
		ad.POST("/lock", admin.LockUser)                     //封禁账号
		ad.POST("/unlock", admin.UnlockUser)                 //解封账号
		r.NoRoute(middleware.HandleNotFound)                 //未注册路由
	}
}
