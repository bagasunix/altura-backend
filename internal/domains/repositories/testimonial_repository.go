package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type TestimonialRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Testimonial, error)
	FindAll(ctx context.Context, activeOnly bool) ([]*entities.Testimonial, error)
	Create(ctx context.Context, t *entities.Testimonial) (*entities.Testimonial, error)
	Update(ctx context.Context, t *entities.Testimonial) (*entities.Testimonial, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
