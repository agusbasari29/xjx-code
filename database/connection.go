package database

import (
	"fmt"
	"log"

	"github.com/agusbasari29/xjx-code/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	dbConfig := config.DbConfig()
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", //postgres
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", //mysql
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // postgres
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // mysql
	if err != nil {
		log.Fatalf("Failed to connecting database!")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to close database connection!")
	}
	conn.Close()
}
