package database

import (
	"CWall/app/models"
	"CWall/config/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var PostTab *gorm.DB
var ReportTab *gorm.DB

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Report{},
		&models.Email{},
		&models.Comment{},
		&models.Block{},
		&models.Love{},
	)
	return err
}
func InitDB() {
	user := config.Config.GetString("database.user")
	pass := config.Config.GetString("database.password")
	host := config.Config.GetString("database.host")
	port := config.Config.GetString("database.port")
	name := config.Config.GetString("database.DBname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	autoMigrate(db)
	DB = db
	PostTab = db.Table("posts")
	ReportTab = db.Table("reports")
}
