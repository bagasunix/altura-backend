package repositories

import (
	"context"

	"github.com/gofrs/uuid"

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
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Property, error)
	FindByName(ctx context.Context, name string) (*entities.Property, error)
	FindAll(ctx context.Context, filter PropertyFilter) ([]*entities.Property, int, error)
	Create(ctx context.Context, product *entities.Property) (*entities.Property, error)
	Update(ctx context.Context, product *entities.Property) (*entities.Property, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
