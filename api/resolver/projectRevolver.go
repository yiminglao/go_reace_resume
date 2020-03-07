package resolver

import (
	"context"

	"github.com/yiminglao/resume_web/api/models"
	"github.com/yiminglao/resume_web/api/repository"
)

func (r *queryResolver) GetProjectByID(ctx context.Context, id int) (*models.Project, error) {
	var err error
	return repository.GetProjectById(id), err
}
func (r *queryResolver) GetAllProject(ctx context.Context) ([]*models.Project, error) {
	var err error
	return repository.GetAllProject(), err
}
