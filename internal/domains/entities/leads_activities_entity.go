package entities

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type LeadActivity struct {
	ID        uuid.UUID
	LeadID    string
	AgentID   *string
	Type      string
	Body      string
	Metadata  json.RawMessage
	IsAuto    bool
	CreatedAt time.Time
}
