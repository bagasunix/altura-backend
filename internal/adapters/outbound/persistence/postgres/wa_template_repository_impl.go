package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type waTemplateRepository struct {
	db *gorm.DB
}

func NewWATemplateRepository(db *gorm.DB) repositories.WATemplateRepository {
	return &waTemplateRepository{db: db}
}

func (r *waTemplateRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.WATemplate, error) {
	panic("unimplemented")
}

func (r *waTemplateRepository) FindByTriggerEvent(ctx context.Context, event string) ([]*entities.WATemplate, error) {
	panic("unimplemented")
}

func (r *waTemplateRepository) FindAll(ctx context.Context, activeOnly bool) ([]*entities.WATemplate, error) {
	panic("unimplemented")
}

func (r *waTemplateRepository) Create(ctx context.Context, tmpl *entities.WATemplate) (*entities.WATemplate, error) {
	panic("unimplemented")
}

func (r *waTemplateRepository) Update(ctx context.Context, tmpl *entities.WATemplate) (*entities.WATemplate, error) {
	panic("unimplemented")
}

func (r *waTemplateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
