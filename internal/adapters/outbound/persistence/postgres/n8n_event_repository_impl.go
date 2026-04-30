package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type n8nEventRepository struct {
	db *gorm.DB
}

func NewN8nEventRepository(db *gorm.DB) repositories.N8nEventRepository {
	return &n8nEventRepository{db: db}
}

func (r *n8nEventRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.N8nEvent, error) {
	panic("unimplemented")
}

func (r *n8nEventRepository) FindAll(ctx context.Context, filter repositories.N8nEventFilter) ([]*entities.N8nEvent, int, error) {
	panic("unimplemented")
}

func (r *n8nEventRepository) FindPending(ctx context.Context, limit int) ([]*entities.N8nEvent, error) {
	panic("unimplemented")
}

func (r *n8nEventRepository) Create(ctx context.Context, event *entities.N8nEvent) (*entities.N8nEvent, error) {
	panic("unimplemented")
}

func (r *n8nEventRepository) Update(ctx context.Context, event *entities.N8nEvent) (*entities.N8nEvent, error) {
	panic("unimplemented")
}

func (r *n8nEventRepository) MarkDone(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

func (r *n8nEventRepository) MarkFailed(ctx context.Context, id uuid.UUID, errMsg string) error {
	panic("unimplemented")
}
