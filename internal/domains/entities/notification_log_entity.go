package entities

import (
	"time"

	"github.com/gofrs/uuid"

	"altura-property/shared/constants"
)

type NotificationLog struct {
	ID           uuid.UUID
	LeadID       *uuid.UUID
	AgentID      *uuid.UUID
	TemplateID   *uuid.UUID
	Channel      constants.NotificationChannel
	Recipient    string
	Message      string
	Status       constants.NotificationStatus
	ExternalID   string
	SentAt       *time.Time
	DeliveredAt  *time.Time
	ErrorMessage string
	CreatedAt    time.Time
}
