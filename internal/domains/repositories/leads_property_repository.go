package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type LeadsPropertyFilter struct {
	Search    string
	Type      constants.TypeProperty
	Status    constants.LeadStatus
	MinPrice  *float64
	MaxPrice  *float64
	Page      int
	PageSize  int
	SortBy    string
	SortOrder string
}

type LeadPropertyRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Lead, error)
	FindByName(ctx context.Context, name string) (*entities.Lead, error)
	FindAll(ctx context.Context, filter LeadsPropertyFilter) ([]*entities.Lead, int, error)
	Create(ctx context.Context, product *entities.Lead) (*entities.Lead, error)
	Update(ctx context.Context, product *entities.Lead) (*entities.Lead, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
