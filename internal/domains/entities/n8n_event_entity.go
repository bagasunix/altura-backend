package entities

import (
	"encoding/json"
	"time"
)

type N8nEvent struct {
	ID           string
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
