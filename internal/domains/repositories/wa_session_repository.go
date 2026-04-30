package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type WASessionFilter struct {
	Status   constants.WASessionStatus
	Page     int
	PageSize int
}

type WASessionRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.WASession, error)
	FindByAgentID(ctx context.Context, agentID uuid.UUID) (*entities.WASession, error)
	FindByJID(ctx context.Context, jid string) (*entities.WASession, error)
	FindAll(ctx context.Context, filter WASessionFilter) ([]*entities.WASession, int, error)
	Create(ctx context.Context, session *entities.WASession) (*entities.WASession, error)
	Update(ctx context.Context, session *entities.WASession) (*entities.WASession, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
