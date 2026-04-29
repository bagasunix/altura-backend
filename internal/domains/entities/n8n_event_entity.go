package entities

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type N8nEvent struct {
	ID           uuid.UUID
	EventType    string
	EntityType   string
	EntityID     *string
	Payload      json.RawMessage
	Status       string
	Attempts     int
	MaxAttempts  int
	ScheduledAt  time.Time
	ProcessedAt  *time.Time
	ErrorMessage string
	CreatedAt    time.Time
}
