package entities

import (
	"encoding/json"
	"time"
)

type WAMessage struct {
	ID        string
	SessionID string
	LeadID    *string

	WAMessageID string
	Direction   string

	FromJID string
	ToJID   string
	ChatJID string

	MessageType string
	Body        string

	MediaURL      string
	MediaMime     string
	MediaSize     int64
	MediaFilename string

	Caption    string
	QuotedID   string
	QuotedBody string

	Status  string
	IsRead  bool
	IsAuto  bool
	IsGroup bool

	SentAt      time.Time
	DeliveredAt *time.Time
	ReadAt      *time.Time

	Metadata  json.RawMessage
	CreatedAt time.Time
}
