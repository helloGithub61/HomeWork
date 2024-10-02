package services

import (
	"CWall/app/models"
	"CWall/config/database"
)

func GetUserNumByEmail(email string) (int, error) {
	var user []models.User
	err := database.DB.Table("users").Where("email =?", email).Find(&user).Error
	return len(user), err
}
func GetUserByAccount(account string) (*models.User, error) {
	var user models.User
	result := database.DB.Table("users").Where("account =?", account).First(&user)
	return &user, result.Error
}

func CreateUser(account string, name string, email string, password string, key string) error {
	result := database.DB.Table("users").Create(&models.User{
		Account:  account,
		Name:     name,
		Password: password,
		Email:    email,
		TKey:     key,
	})
	return result.Error
}
func UpdateUserPassword(account string, password string) error {
	user, _ := GetUserByAccount(account)
	user.Password = password
	result := database.DB.Table("users").Save(&user)
	return result.Error
}
func UpdateUserName(account string, name string) error {
	user, _ := GetUserByAccount(account)
	user.Name = name
	result := database.DB.Table("users").Save(&user)
	return result.Error
}
func BlockUser(account string, baccount string) error {
	block := models.Block{Account: account, Block: baccount}
	err := database.DB.Table("blocks").Create(&block).Error
	return err
}
func GetAllUserList() (userList []models.User, err error) {
	err = database.DB.Table("users").Find(&userList).Error
	return userList, err
}
func UpdateUserStatus(acc string, status int) (err error) {
	user, _ := GetUserByAccount(acc)
	user.Status = status
	err = database.DB.Table("users").Save(&user).Error
	return err
}
func UpdateUserImage(acc string, src string) (err error) {
	user, _ := GetUserByAccount(acc)
	user.Image = src
	err = database.DB.Table("users").Save(&user).Error
	return err
}
