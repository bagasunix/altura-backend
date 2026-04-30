package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type waAutoReplyRepository struct {
	db *gorm.DB
}

func NewWAAutoReplyRepository(db *gorm.DB) repositories.WAAutoReplyRepository {
	return &waAutoReplyRepository{db: db}
}

func (r *waAutoReplyRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.WAAutoReply, error) {
	panic("unimplemented")
}

func (r *waAutoReplyRepository) FindActive(ctx context.Context, sessionID *uuid.UUID) ([]*entities.WAAutoReply, error) {
	panic("unimplemented")
}

func (r *waAutoReplyRepository) FindAll(ctx context.Context) ([]*entities.WAAutoReply, error) {
	panic("unimplemented")
}

func (r *waAutoReplyRepository) Create(ctx context.Context, r2 *entities.WAAutoReply) (*entities.WAAutoReply, error) {
	panic("unimplemented")
}

func (r *waAutoReplyRepository) Update(ctx context.Context, r2 *entities.WAAutoReply) (*entities.WAAutoReply, error) {
	panic("unimplemented")
}

func (r *waAutoReplyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
