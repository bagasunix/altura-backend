package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type leadNoteRepository struct {
	db *gorm.DB
}

func NewLeadNoteRepository(db *gorm.DB) repositories.LeadNoteRepository {
	return &leadNoteRepository{db: db}
}

func (r *leadNoteRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.LeadNote, error) {
	panic("unimplemented")
}

func (r *leadNoteRepository) FindByLeadID(ctx context.Context, leadID uuid.UUID) ([]*entities.LeadNote, error) {
	panic("unimplemented")
}

func (r *leadNoteRepository) Create(ctx context.Context, note *entities.LeadNote) (*entities.LeadNote, error) {
	panic("unimplemented")
}

func (r *leadNoteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
