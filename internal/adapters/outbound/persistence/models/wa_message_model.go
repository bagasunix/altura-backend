package models

import "time"

type WAMessage struct {
	BaseModel
	SessionID     string `gorm:"type:uuid;not null;index"`
	Session       WASession
	LeadID        *string `gorm:"type:uuid;index"`
	Lead          *Lead
	WAMessageID   string `gorm:"type:text;not null"`
	Direction     string `gorm:"type:wa_message_direction;not null;index"`
	FromJID       string `gorm:"type:text;not null"`
	ToJID         string `gorm:"type:text;not null"`
	ChatJID       string `gorm:"type:text;not null;index"`
	MessageType   string `gorm:"type:wa_message_type;not null;default:'text'"`
	Body          string `gorm:"type:text;not null;default:''"`
	MediaURL      string `gorm:"type:text"`
	MediaMime     string `gorm:"type:text"`
	MediaSize     int64
	MediaFilename string                 `gorm:"type:text"`
	Caption       string                 `gorm:"type:text"`
	QuotedID      string                 `gorm:"type:text"`
	QuotedBody    string                 `gorm:"type:text"`
	Status        string                 `gorm:"type:wa_message_status;not null;default:'sent'"`
	IsRead        bool                   `gorm:"not null;default:false;index"`
	IsAuto        bool                   `gorm:"not null;default:false"`
	IsGroup       bool                   `gorm:"not null;default:false"`
	SentAt        time.Time              `gorm:"type:timestamptz;not null;default:now()"`
	DeliveredAt   *time.Time             `gorm:"type:timestamptz"`
	ReadAt        *time.Time             `gorm:"type:timestamptz"`
	Metadata      map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
}
