package resolver

import (
	"context"

	"github.com/yiminglao/resume_web/api/models"
	"github.com/yiminglao/resume_web/api/repository"
)

func (r *queryResolver) GetExperienceByID(ctx context.Context, id int) (*models.Experience, error) {
	var err error
	return repository.GetExperienceById(id), err
}

func (r *queryResolver) GetAllExp(ctx context.Context) ([]*models.Experience, error) {
	var err error
	return repository.GetAllExp(), err
}
