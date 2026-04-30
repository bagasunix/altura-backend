package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type leadRepository struct {
	db *gorm.DB
}

func NewLeadRepository(db *gorm.DB) repositories.LeadPropertyRepository {
	return &leadRepository{db: db}
}

func (r *leadRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Lead, error) {
	panic("unimplemented")
}

func (r *leadRepository) FindByName(ctx context.Context, name string) (*entities.Lead, error) {
	panic("unimplemented")
}

func (r *leadRepository) FindAll(ctx context.Context, filter repositories.LeadsPropertyFilter) ([]*entities.Lead, int, error) {
	panic("unimplemented")
}

func (r *leadRepository) Create(ctx context.Context, lead *entities.Lead) (*entities.Lead, error) {
	panic("unimplemented")
}

func (r *leadRepository) Update(ctx context.Context, lead *entities.Lead) (*entities.Lead, error) {
	panic("unimplemented")
}

func (r *leadRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
