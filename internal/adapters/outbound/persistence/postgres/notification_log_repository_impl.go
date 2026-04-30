package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type notificationLogRepository struct {
	db *gorm.DB
}

func NewNotificationLogRepository(db *gorm.DB) repositories.NotificationLogRepository {
	return &notificationLogRepository{db: db}
}

func (r *notificationLogRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.NotificationLog, error) {
	panic("unimplemented")
}

func (r *notificationLogRepository) FindByLeadID(ctx context.Context, leadID uuid.UUID, page, pageSize int) ([]*entities.NotificationLog, int, error) {
	panic("unimplemented")
}

func (r *notificationLogRepository) FindAll(ctx context.Context, filter repositories.NotificationLogFilter) ([]*entities.NotificationLog, int, error) {
	panic("unimplemented")
}

func (r *notificationLogRepository) Create(ctx context.Context, log *entities.NotificationLog) (*entities.NotificationLog, error) {
	panic("unimplemented")
}

func (r *notificationLogRepository) Update(ctx context.Context, log *entities.NotificationLog) (*entities.NotificationLog, error) {
	panic("unimplemented")
}
