package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type waSessionRepository struct {
	db *gorm.DB
}

func NewWASessionRepository(db *gorm.DB) repositories.WASessionRepository {
	return &waSessionRepository{db: db}
}

func (r *waSessionRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.WASession, error) {
	panic("unimplemented")
}

func (r *waSessionRepository) FindByAgentID(ctx context.Context, agentID uuid.UUID) (*entities.WASession, error) {
	panic("unimplemented")
}

func (r *waSessionRepository) FindByJID(ctx context.Context, jid string) (*entities.WASession, error) {
	panic("unimplemented")
}

func (r *waSessionRepository) FindAll(ctx context.Context, filter repositories.WASessionFilter) ([]*entities.WASession, int, error) {
	panic("unimplemented")
}

func (r *waSessionRepository) Create(ctx context.Context, session *entities.WASession) (*entities.WASession, error) {
	panic("unimplemented")
}

func (r *waSessionRepository) Update(ctx context.Context, session *entities.WASession) (*entities.WASession, error) {
	panic("unimplemented")
}

func (r *waSessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
