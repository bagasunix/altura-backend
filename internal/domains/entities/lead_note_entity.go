package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type LeadNote struct {
	ID        uuid.UUID
	LeadID    uuid.UUID
	Body      string
	CreatedAt time.Time
}
