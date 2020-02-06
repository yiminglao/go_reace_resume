package resolver

import (
	"context"

	"github.com/yiminglao/resume_web/api/models"
)

func (r *queryResolver) GetMessageByID(ctx context.Context, id int) (*models.Message, error) {
	panic("not implemented")
}
func (r *queryResolver) GetAllMessage(ctx context.Context) ([]*models.Message, error) {
	panic("not implemented")
}
