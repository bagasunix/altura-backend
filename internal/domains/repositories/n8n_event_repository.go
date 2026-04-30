package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type N8nEventFilter struct {
	Status     constants.N8nEventStatus
	EventType  string
	EntityType string
	EntityID   *uuid.UUID
	Page       int
	PageSize   int
}

type N8nEventRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.N8nEvent, error)
	FindAll(ctx context.Context, filter N8nEventFilter) ([]*entities.N8nEvent, int, error)
	FindPending(ctx context.Context, limit int) ([]*entities.N8nEvent, error)
	Create(ctx context.Context, event *entities.N8nEvent) (*entities.N8nEvent, error)
	Update(ctx context.Context, event *entities.N8nEvent) (*entities.N8nEvent, error)
	MarkDone(ctx context.Context, id uuid.UUID) error
	MarkFailed(ctx context.Context, id uuid.UUID, errMsg string) error
}
