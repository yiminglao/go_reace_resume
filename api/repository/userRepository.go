package repository

import (
	"github.com/yiminglao/resume_web/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Get User
func GetUser() models.User {
	var user models.User
	db.Model(&models.User{}).First(&user)
	return user
}
