package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type FAQRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.FAQ, error)
	FindAll(ctx context.Context, activeOnly bool) ([]*entities.FAQ, error)
	Create(ctx context.Context, faq *entities.FAQ) (*entities.FAQ, error)
	Update(ctx context.Context, faq *entities.FAQ) (*entities.FAQ, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
