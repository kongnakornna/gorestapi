package items

import (
	"context"

	"github.com/google/uuid"
	"github.com/kongnakornna/gorestapi/internal"
	"github.com/kongnakornna/gorestapi/internal/models"
)

type ItemUseCaseI interface {
	internal.UseCaseI[models.Item]
	GetMultiByOwnerId(ctx context.Context, ownerId uuid.UUID, limit, offset int) ([]*models.Item, error)
	CreateWithOwner(ctx context.Context, ownerId uuid.UUID, exp *models.Item) (*models.Item, error)
	DeleteWithoutGet(ctx context.Context, id uuid.UUID) error
}
