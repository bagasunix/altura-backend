package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type leadActivityRepository struct {
	db *gorm.DB
}

func NewLeadActivityRepository(db *gorm.DB) repositories.LeadActivityRepository {
	return &leadActivityRepository{db: db}
}

func (r *leadActivityRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.LeadActivity, error) {
	panic("unimplemented")
}

func (r *leadActivityRepository) FindByLeadID(ctx context.Context, leadID uuid.UUID, page, pageSize int) ([]*entities.LeadActivity, int, error) {
	panic("unimplemented")
}

func (r *leadActivityRepository) FindAll(ctx context.Context, filter repositories.LeadActivityFilter) ([]*entities.LeadActivity, int, error) {
	panic("unimplemented")
}

func (r *leadActivityRepository) Create(ctx context.Context, activity *entities.LeadActivity) (*entities.LeadActivity, error) {
	panic("unimplemented")
}

func (r *leadActivityRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
