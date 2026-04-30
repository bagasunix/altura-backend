package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type waMessageRepository struct {
	db *gorm.DB
}

func NewWAMessageRepository(db *gorm.DB) repositories.WAMessageRepository {
	return &waMessageRepository{db: db}
}

func (r *waMessageRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.WAMessage, error) {
	panic("unimplemented")
}

func (r *waMessageRepository) FindByChatJID(ctx context.Context, sessionID uuid.UUID, chatJID string, page, pageSize int) ([]*entities.WAMessage, int, error) {
	panic("unimplemented")
}

func (r *waMessageRepository) FindByLeadID(ctx context.Context, leadID uuid.UUID, page, pageSize int) ([]*entities.WAMessage, int, error) {
	panic("unimplemented")
}

func (r *waMessageRepository) FindUnread(ctx context.Context, sessionID uuid.UUID) ([]*entities.WAMessage, error) {
	panic("unimplemented")
}

func (r *waMessageRepository) FindAll(ctx context.Context, filter repositories.WAMessageFilter) ([]*entities.WAMessage, int, error) {
	panic("unimplemented")
}

func (r *waMessageRepository) Create(ctx context.Context, msg *entities.WAMessage) (*entities.WAMessage, error) {
	panic("unimplemented")
}

func (r *waMessageRepository) Update(ctx context.Context, msg *entities.WAMessage) (*entities.WAMessage, error) {
	panic("unimplemented")
}
