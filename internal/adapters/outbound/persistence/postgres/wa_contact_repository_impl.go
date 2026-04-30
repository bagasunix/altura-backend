package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type waContactRepository struct {
	db *gorm.DB
}

func NewWAContactRepository(db *gorm.DB) repositories.WAContactRepository {
	return &waContactRepository{db: db}
}

func (r *waContactRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.WAContact, error) {
	panic("unimplemented")
}

func (r *waContactRepository) FindByJID(ctx context.Context, sessionID uuid.UUID, jid string) (*entities.WAContact, error) {
	panic("unimplemented")
}

func (r *waContactRepository) FindByPhone(ctx context.Context, sessionID uuid.UUID, phone string) (*entities.WAContact, error) {
	panic("unimplemented")
}

func (r *waContactRepository) FindByLeadID(ctx context.Context, leadID uuid.UUID) ([]*entities.WAContact, error) {
	panic("unimplemented")
}

func (r *waContactRepository) FindAll(ctx context.Context, sessionID uuid.UUID, page, pageSize int) ([]*entities.WAContact, int, error) {
	panic("unimplemented")
}

func (r *waContactRepository) Upsert(ctx context.Context, contact *entities.WAContact) (*entities.WAContact, error) {
	panic("unimplemented")
}

func (r *waContactRepository) Update(ctx context.Context, contact *entities.WAContact) (*entities.WAContact, error) {
	panic("unimplemented")
}

func (r *waContactRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
