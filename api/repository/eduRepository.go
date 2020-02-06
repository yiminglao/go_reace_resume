package repository

import (
	"github.com/yiminglao/resume_web/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Get GetEduById
func GetEduById(id int) *models.Edu {
	var edu models.Edu
	db.Model(&models.Edu{}).Where("id = ?", id).First(&edu)
	return &edu
}

//Create new edu row
func AddEdu(edu *models.Edu) *models.Edu {
	db.Model(&models.Edu{}).Create(edu)
	return edu
}

func UpdateEdu(edu *models.Edu) *models.Edu {
	db.Model(&models.Edu{}).Update(edu)
	return edu
}

//Get all edus
func GetAllEdu() []*models.Edu {
	var edus []*models.Edu
	db.Model(&models.Edu{}).Find(&edus)
	return edus
}

func DelEdu(id int) string {
	db.Model(&models.Edu{}).Delete(&models.Edu{ID: id})
	return "deleted"
}
