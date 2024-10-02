package main

import (
	"CWall/config/database"
	"CWall/config/router"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)
	database.InitDB()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
