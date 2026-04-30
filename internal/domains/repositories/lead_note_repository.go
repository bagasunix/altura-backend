package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type LeadNoteRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.LeadNote, error)
	FindByLeadID(ctx context.Context, leadID uuid.UUID) ([]*entities.LeadNote, error)
	Create(ctx context.Context, note *entities.LeadNote) (*entities.LeadNote, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
