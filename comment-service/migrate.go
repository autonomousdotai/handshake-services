package main

import (
	"github.com/jinzhu/gorm"
	"github.com/ninjadotorg/handshake-services/comment-service/models"
)

func main() {

	//
	var db *gorm.DB = models.Database()
	defer db.Close()

	db.AutoMigrate(&models.Comment{})
}
