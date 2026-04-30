package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type WAAutoReplyRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.WAAutoReply, error)
	FindActive(ctx context.Context, sessionID *uuid.UUID) ([]*entities.WAAutoReply, error)
	FindAll(ctx context.Context) ([]*entities.WAAutoReply, error)
	Create(ctx context.Context, r *entities.WAAutoReply) (*entities.WAAutoReply, error)
	Update(ctx context.Context, r *entities.WAAutoReply) (*entities.WAAutoReply, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
