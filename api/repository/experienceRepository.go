package repository

import (
	"github.com/yiminglao/resume_web/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Get getExperienceById
func GetExperienceById(id int) *models.Experience {
	var exp models.Experience
	db.Model(&models.Experience{}).Where(&models.Experience{ID: id}).First(&exp)
	return &exp
}

func GetAllExp() []*models.Experience {
	var exps []*models.Experience
	db.Model(&models.Experience{}).Find(&exps)
	return exps
}
