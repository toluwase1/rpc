package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"rpc-client-server/models"
)

var Db *gorm.DB

func InitPostgresDB(dbUser string, dbPwd string, dbName string, dbHost string, dbPort string) {
	var err error
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPwd)
	Db, err = gorm.Open("postgres", dbUrl)
	if err != nil {
		fmt.Printf("Failed to connect to database %v", err)
		log.Panic("Db connection failed")
	} else {
		fmt.Println("db connected successfully")
		Automigrate()
	}
}

func Automigrate() {
	Db.AutoMigrate(&models.Authorization{}, &models.Subscription{})
}
