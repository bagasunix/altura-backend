package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type WAContactRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.WAContact, error)
	FindByJID(ctx context.Context, sessionID uuid.UUID, jid string) (*entities.WAContact, error)
	FindByPhone(ctx context.Context, sessionID uuid.UUID, phone string) (*entities.WAContact, error)
	FindByLeadID(ctx context.Context, leadID uuid.UUID) ([]*entities.WAContact, error)
	FindAll(ctx context.Context, sessionID uuid.UUID, page, pageSize int) ([]*entities.WAContact, int, error)
	Upsert(ctx context.Context, contact *entities.WAContact) (*entities.WAContact, error)
	Update(ctx context.Context, contact *entities.WAContact) (*entities.WAContact, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
