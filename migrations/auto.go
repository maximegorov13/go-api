package main

import (
	"github.com/joho/godotenv"
	"github.com/maximegorov13/go-api/internal/link"
	"github.com/maximegorov13/go-api/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&link.Link{},
		&user.User{},
	)
}
