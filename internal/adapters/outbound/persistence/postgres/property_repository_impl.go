package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/adapters/outbound/persistence/models"
	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type gormProvider struct {
	db *gorm.DB
}

// Create implements [repositories.PropertyRepository].
func (g *gormProvider) Create(ctx context.Context, product *entities.Property) (result models.SingleResult[*entities.Property]) {
	panic("unimplemented")

}

// Delete implements [repositories.PropertyRepository].
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// FindAll implements [repositories.PropertyRepository].
func (g *gormProvider) FindAll(ctx context.Context, filter repositories.PropertyFilter) (result models.SliceResult[*entities.Property]) {
	panic("unimplemented")
}

// FindByID implements [repositories.PropertyRepository].
func (g *gormProvider) FindByID(ctx context.Context, id uuid.UUID) (result models.SingleResult[*entities.Property]) {
	panic("unimplemented")
}

// FindByName implements [repositories.PropertyRepository].
func (g *gormProvider) FindByName(ctx context.Context, name string) (result models.SingleResult[*entities.Property]) {
	panic("unimplemented")
}

// Update implements [repositories.PropertyRepository].
func (g *gormProvider) Update(ctx context.Context, product *entities.Property) (result models.SingleResult[*entities.Property]) {
	panic("unimplemented")
}

func NewProductRepository(db *gorm.DB) repositories.PropertyRepository {
	return &gormProvider{db: db}
}
