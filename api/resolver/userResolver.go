package resolver

import (
	"context"

	"github.com/yiminglao/resume_web/api/models"
	"github.com/yiminglao/resume_web/api/repository"
)

func (r *queryResolver) GetUser(ctx context.Context) (*models.User, error) {
	var err error
	var user models.User
	user = repository.GetUser()

	return &user, err
}
