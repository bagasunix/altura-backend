package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type testimonialRepository struct {
	db *gorm.DB
}

func NewTestimonialRepository(db *gorm.DB) repositories.TestimonialRepository {
	return &testimonialRepository{db: db}
}

func (r *testimonialRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Testimonial, error) {
	panic("unimplemented")
}

func (r *testimonialRepository) FindAll(ctx context.Context, activeOnly bool) ([]*entities.Testimonial, error) {
	panic("unimplemented")
}

func (r *testimonialRepository) Create(ctx context.Context, t *entities.Testimonial) (*entities.Testimonial, error) {
	panic("unimplemented")
}

func (r *testimonialRepository) Update(ctx context.Context, t *entities.Testimonial) (*entities.Testimonial, error) {
	panic("unimplemented")
}

func (r *testimonialRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
