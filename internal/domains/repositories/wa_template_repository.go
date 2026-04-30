package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type WATemplateRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.WATemplate, error)
	FindByTriggerEvent(ctx context.Context, event string) ([]*entities.WATemplate, error)
	FindAll(ctx context.Context, activeOnly bool) ([]*entities.WATemplate, error)
	Create(ctx context.Context, tmpl *entities.WATemplate) (*entities.WATemplate, error)
	Update(ctx context.Context, tmpl *entities.WATemplate) (*entities.WATemplate, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
