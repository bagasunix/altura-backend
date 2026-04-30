package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type WAMessageFilter struct {
	SessionID  *uuid.UUID
	LeadID     *uuid.UUID
	ChatJID    string
	Direction  constants.WAMessageDirection
	IsRead     *bool
	Page       int
	PageSize   int
	SortOrder  string
}

type WAMessageRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.WAMessage, error)
	FindByChatJID(ctx context.Context, sessionID uuid.UUID, chatJID string, page, pageSize int) ([]*entities.WAMessage, int, error)
	FindByLeadID(ctx context.Context, leadID uuid.UUID, page, pageSize int) ([]*entities.WAMessage, int, error)
	FindUnread(ctx context.Context, sessionID uuid.UUID) ([]*entities.WAMessage, error)
	FindAll(ctx context.Context, filter WAMessageFilter) ([]*entities.WAMessage, int, error)
	Create(ctx context.Context, msg *entities.WAMessage) (*entities.WAMessage, error)
	Update(ctx context.Context, msg *entities.WAMessage) (*entities.WAMessage, error)
}
