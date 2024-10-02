package services

import (
	"CWall/app/models"
	"CWall/config/database"
	"time"
)

func GetPostByID(pid int) (post *models.Post, err error) {
	err = database.PostTab.First(&post, pid).Error
	return post, err
}

func CreatePost(account string, Content string, ShowType int) error {
	result := database.PostTab.Create(&models.Post{
		CreateTime:  time.Now(),
		Account:     account,
		PostContent: Content,
		ShowType:    ShowType,
	})
	return result.Error
}
func GetAllPostList() (postList []models.Post, err error) {
	err = database.PostTab.Find(&postList).Error
	return postList, err
}
func GetHomePostList(p int) (postLst []models.Post, err error) {
	err = database.PostTab.Offset(20 * p).Limit(20).Find(&postLst).Error
	return postLst, err
}
func GetAdminReviewList(p int) (reportLst []models.Report, err error) {
	err = database.ReportTab.Offset(20 * p).Limit(20).Find(&reportLst).Error
	return reportLst, err
}
func GetPostListByAccount(account string) (postLst []models.Post, err error) {
	err = database.PostTab.Where("account =?", account).Find(&postLst).Error
	return postLst, err
}

func UpdatePostContent(pid int, content string) error {
	post, _ := GetPostByID(pid)
	post.PostContent = content
	post.CreateTime = time.Now()
	err := database.PostTab.Where("id =?", pid).Save(&post).Error
	return err
}

func DeletePost(pid int) (err error) {
	err = database.PostTab.Delete(&models.Post{}, pid).Error
	return err
}

func ReportPost(account string, pid int, reason string) (err error) {
	err = database.ReportTab.Create(&models.Report{
		PostID:  pid,
		Account: account,
		Reason:  reason,
	}).Error
	return err
}

func GetReportList(account string) (reportList []models.Report, err error) {
	err = database.ReportTab.Where("account=? ", account).Find(&reportList).Error
	return reportList, err
}

func GetReportByPostID(pid int) (report *models.Report, err error) {
	err = database.DB.Table("reports").Where("post_id=?", pid).First(&report).Error
	return report, err
}

func UpdateReportStatus(pid int, approval int) (err error) {
	var report models.Report
	database.ReportTab.Where("post_id=?", pid).First(&report)
	report.Status = approval
	err = database.ReportTab.Where("post_id=?", pid).Save(&report).Error
	return err
}
func UpdateReportContent(pid int) (err error) {
	var post models.Post
	var report models.Report
	database.ReportTab.Where("post_id=?", pid).First(&report)
	database.PostTab.First(&post, pid)
	report.Content = post.PostContent
	err = database.ReportTab.Where("post_id=?", pid).Save(&report).Error
	return err
}
func GetLoveByAccAndPID(account string, pid int) (love *models.Love, err error) {
	err = database.DB.Table("loves").Where("account=? AND post_id=?", account, pid).First(&love).Error
	return love, err
}
func UpdateLike(love models.Love) (err error) {
	err = database.DB.Table("loves").Save(&love).Error
	return err
}
func UpdateCollect(love models.Love) (err error) {
	err = database.DB.Table("loves").Save(&love).Error
	return err
}
func CreateLove(love models.Love) (err error) {
	err = database.DB.Table("loves").Create(&love).Error
	return err
}
func GetLoveList(account string) (loveList []models.Love, err error) {
	err = database.DB.Table("loves").Where("account=?", account).Find(&loveList).Error
	return loveList, err
}

func GetAllLoveList() (loveList []models.Love, err error) {
	err = database.DB.Table("loves").Find(&loveList).Error
	return loveList, err
}
