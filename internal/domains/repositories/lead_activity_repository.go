package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
	"altura-property/shared/constants"
)

type LeadActivityFilter struct {
	LeadID   *uuid.UUID
	AgentID  *uuid.UUID
	Type     constants.ActivityType
	IsAuto   *bool
	Page     int
	PageSize int
}

type LeadActivityRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.LeadActivity, error)
	FindByLeadID(ctx context.Context, leadID uuid.UUID, page, pageSize int) ([]*entities.LeadActivity, int, error)
	FindAll(ctx context.Context, filter LeadActivityFilter) ([]*entities.LeadActivity, int, error)
	Create(ctx context.Context, activity *entities.LeadActivity) (*entities.LeadActivity, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
