package repository

import (
	"github.com/yiminglao/resume_web/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Get GetMessageById
func GetMessageById(id int) models.Message {
	var message models.Message
	db.Model(&models.Message{}).Where("id = ", id).First(&message)
	return message
}
