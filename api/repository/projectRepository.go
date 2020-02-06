package repository

import (
	"github.com/yiminglao/resume_web/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Get GetProjectById
func GetProjectById(id int) *models.Project {
	var project models.Project
	db.Model(&models.Project{}).Where(&models.Project{ID: id}).First(&project)
	return &project
}

func GetAllProject() []*models.Project {
	var projects []*models.Project
	db.Model(&models.Project{}).Find(&projects)
	return projects
}
