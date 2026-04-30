package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type NotificationLogFilter struct {
	LeadID  *uuid.UUID
	AgentID *uuid.UUID
	Channel constants.NotificationChannel
	Status  constants.NotificationStatus
	Page    int
	PageSize int
}

type NotificationLogRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.NotificationLog, error)
	FindByLeadID(ctx context.Context, leadID uuid.UUID, page, pageSize int) ([]*entities.NotificationLog, int, error)
	FindAll(ctx context.Context, filter NotificationLogFilter) ([]*entities.NotificationLog, int, error)
	Create(ctx context.Context, log *entities.NotificationLog) (*entities.NotificationLog, error)
	Update(ctx context.Context, log *entities.NotificationLog) (*entities.NotificationLog, error)
}
