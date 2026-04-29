package models

import "time"

type WASession struct {
	BaseModel
	AgentID      string `gorm:"type:uuid;not null;unique"`
	Agent        Agent
	JID          string     `gorm:"type:text;unique"`
	Phone        string     `gorm:"type:text;not null;default:''"`
	DisplayName  string     `gorm:"type:text;not null;default:''"`
	Status       string     `gorm:"type:wa_session_status;not null;default:'disconnected';index"`
	QRCode       string     `gorm:"type:text"`
	QRExpiresAt  *time.Time `gorm:"type:timestamptz"`
	Platform     string     `gorm:"type:text;default:''"`
	BatteryPct   int
	IsOnline     bool       `gorm:"not null;default:false"`
	LastSeenAt   *time.Time `gorm:"type:timestamptz"`
	MessagesSent int        `gorm:"not null;default:0"`
	MessagesRecv int        `gorm:"not null;default:0"`
}
