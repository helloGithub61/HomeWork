package services

import (
	"CWall/app/models"
	"CWall/config/database"
	"time"
)

func GetCommentByID(pid int) (post *models.Comment, err error) {
	err = database.DB.Table("comments").First(&post, pid).Error
	return post, err
}
func CreateComment(account string, Content string, PostID int, Target int) error {
	result := database.DB.Table("comments").Create(&models.Comment{
		CreateTime: time.Now(),
		Account:    account,
		ComContent: Content,
		Target:     Target,
		PostID:     PostID,
	})
	return result.Error
}
func GetCommentByPostID(pid int) ([]models.Comment, error) {
	var commentLst []models.Comment
	err := database.DB.Table("comments").Where("post_id=?", pid).Find(&commentLst).Error
	return commentLst, err
}
