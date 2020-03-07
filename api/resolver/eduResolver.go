package resolver

import (
	"context"

	"github.com/yiminglao/resume_web/api/models"
	"github.com/yiminglao/resume_web/api/repository"
)

func (r *queryResolver) GetEduByID(ctx context.Context, id int) (*models.Edu, error) {
	var err error
	var edu *models.Edu
	edu = repository.GetEduById(id)
	return edu, err
}

func (r *queryResolver) GetAllEdu(ctx context.Context) ([]*models.Edu, error) {
	var err error
	return repository.GetAllEdu(), err
}

func (r *mutationResolver) AddEdu(ctx context.Context, input *models.EduInput) (*models.Edu, error) {
	edu := models.Edu{
		School:         input.School,
		Major:          input.Major,
		Gpa:            input.Gpa,
		GraduationDate: input.GraduationDate,
		Description:    input.Description,
	}
	return repository.AddEdu(&edu), nil
}
func (r *mutationResolver) UpdateEdu(ctx context.Context, input *models.EduInput) (*models.Edu, error) {
	edu := models.Edu{
		ID:             *input.ID,
		School:         input.School,
		Major:          input.Major,
		Gpa:            input.Gpa,
		GraduationDate: input.GraduationDate,
		Description:    input.Description,
	}
	return repository.UpdateEdu(&edu), nil
}
func (r *mutationResolver) DelEdu(ctx context.Context, id int) (*string, error) {
	var err error
	result := repository.DelEdu(id)
	return &result, err

}
