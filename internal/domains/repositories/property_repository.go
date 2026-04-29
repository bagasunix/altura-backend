package repositories

import (
	"context"
	"time"

	"github.com/gofrs/uuid"

	"altura-property/internal/adapters/outbound/persistence/models"
	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type PropertyFilter struct {
	Search    string
	Type      constants.TypeProperty
	Status    constants.StatusProperty
	MinPrice  *float64
	MaxPrice  *float64
	Page      int
	PageSize  int
	SortBy    string
	SortOrder string
}

type PropertyRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.Property])
	FindByName(ctx context.Context, name string) (result models.SingleResult[*entities.Property])
	FindAll(ctx context.Context, filter PropertyFilter) (result models.SliceResult[*entities.Property])
	Create(ctx context.Context, product *entities.Property) (result models.SingleResult[*entities.Property])
	Update(ctx context.Context, product *entities.Property) (result models.SingleResult[*entities.Property])
	Delete(ctx context.Context, id uuid.UUID) error
}

type PropertyCachePort interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Property, error)
	SetByID(ctx context.Context, product *entities.Property, ttl time.Duration) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
