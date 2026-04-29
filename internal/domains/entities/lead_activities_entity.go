package entities

import (
	"encoding/json"
	"time"
)

type LeadActivity struct {
	ID        string
	LeadID    string
	AgentID   *string
	Type      string
	Body      string
	Metadata  json.RawMessage
	IsAuto    bool
	CreatedAt time.Time
}
