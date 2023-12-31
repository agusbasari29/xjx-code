package main

import (
	"fmt"
	"os"

	"github.com/agusbasari29/xjx-code/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDBConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate()
	g := gin.Default()
	g.GET("/seeder", func(ctx *gin.Context) {
		fmt.Println("Test seeders function")
	})
	g.Run(os.Getenv("SERVER_PORT"))
}
