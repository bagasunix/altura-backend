package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type agentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) repositories.AgentRepository {
	return &agentRepository{db: db}
}

func (r *agentRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Agent, error) {
	panic("unimplemented")
}

func (r *agentRepository) FindByPhone(ctx context.Context, phone string) (*entities.Agent, error) {
	panic("unimplemented")
}

func (r *agentRepository) FindAll(ctx context.Context, filter repositories.AgentFilter) ([]*entities.Agent, int, error) {
	panic("unimplemented")
}

func (r *agentRepository) Create(ctx context.Context, agent *entities.Agent) (*entities.Agent, error) {
	panic("unimplemented")
}

func (r *agentRepository) Update(ctx context.Context, agent *entities.Agent) (*entities.Agent, error) {
	panic("unimplemented")
}

func (r *agentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
