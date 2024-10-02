package services

import (
	"CWall/app/models"
	"CWall/config/database"
	"time"
)

func CreateEmail(email string, code string) error {
	result := database.DB.Table("emails").Create(&models.Email{
		Create_time: time.Now(),
		Address:     email,
		Code:        code,
	})
	return result.Error
}
func GetEmailByAddress(address string) (email *models.Email, err error) {
	err = database.DB.Table("emails").Where("address=?", address).First(&email).Error
	return email, err
}
func UpdateEmail(address string, code string) (err error) {
	email, _ := GetEmailByAddress(address)
	email.Code = code
	email.Create_time = time.Now()
	err = database.DB.Table("emails").Where("address=?", address).Save(&email).Error
	return err
}
func DeleteEmail(address string) (err error) {
	err = database.DB.Table("emails").Where("address=?", address).Delete(&models.Email{}).Error
	return err
}
