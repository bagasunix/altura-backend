package entities

import (
	"time"

	"altura-property/shared/constants"
)

type WASession struct {
	BaseModel
	AgentID      string
	JID          string
	Phone        string
	DisplayName  string
	Status       constants.WASessionStatus
	QRCode       string
	QRExpiresAt  *time.Time
	Platform     string
	BatteryPct   int
	IsOnline     bool
	LastSeenAt   *time.Time
	MessagesSent int
	MessagesRecv int
}
