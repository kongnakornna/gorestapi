package items

import (
	"context"

	"gorestapi/internal"
	"gorestapi/internal/models"

	"github.com/google/uuid"
)

type ItemPgRepository interface {
	internal.PgRepository[models.Item]
	GetMultiByOwnerId(ctx context.Context, ownerId uuid.UUID, limit, offset int) ([]*models.Item, error)
	CreateWithOwner(ctx context.Context, ownerId uuid.UUID, exp *models.Item) (*models.Item, error)
	DeleteWithoutGet(ctx context.Context, id uuid.UUID) error
}
