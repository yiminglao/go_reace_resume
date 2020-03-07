package resolver

import (
	"context"

	"github.com/yiminglao/resume_web"
	"github.com/yiminglao/resume_web/api/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() resume_web.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() resume_web.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }
