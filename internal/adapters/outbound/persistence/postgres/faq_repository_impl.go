package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type faqRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) repositories.FAQRepository {
	return &faqRepository{db: db}
}

func (r *faqRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.FAQ, error) {
	panic("unimplemented")
}

func (r *faqRepository) FindAll(ctx context.Context, activeOnly bool) ([]*entities.FAQ, error) {
	panic("unimplemented")
}

func (r *faqRepository) Create(ctx context.Context, faq *entities.FAQ) (*entities.FAQ, error) {
	panic("unimplemented")
}

func (r *faqRepository) Update(ctx context.Context, faq *entities.FAQ) (*entities.FAQ, error) {
	panic("unimplemented")
}

func (r *faqRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
