package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConnection struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

func DbConfig() DbConnection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	dbConfig := DbConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	return dbConfig
}
