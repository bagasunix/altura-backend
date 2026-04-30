package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type AgentFilter struct {
	Search    string
	IsActive  *bool
	Page      int
	PageSize  int
	SortBy    string
	SortOrder string
}

type AgentRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Agent, error)
	FindByPhone(ctx context.Context, phone string) (*entities.Agent, error)
	FindAll(ctx context.Context, filter AgentFilter) ([]*entities.Agent, int, error)
	Create(ctx context.Context, agent *entities.Agent) (*entities.Agent, error)
	Update(ctx context.Context, agent *entities.Agent) (*entities.Agent, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
