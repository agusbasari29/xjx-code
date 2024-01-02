package main

import (
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
	gin.SetMode(gin.ReleaseMode)
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Test init project",
		})
	})
	g.Run(os.Getenv("SERVER_PORT"))
}
