package models

import "time"

type NotificationLog struct {
	BaseModel
	LeadID       *string    `gorm:"type:uuid"`
	Lead         *Lead
	AgentID      *string    `gorm:"type:uuid"`
	Agent        *Agent
	TemplateID   *string    `gorm:"type:uuid"`
	Template     *WATemplate
	Channel      string     `gorm:"type:notification_channel;not null;default:'whatsapp';index"`
	Recipient    string     `gorm:"type:text;not null"`
	Message      string     `gorm:"type:text;not null"`
	Status       string     `gorm:"type:notification_status;not null;default:'queued';index"`
	ExternalID   string     `gorm:"type:text"`
	SentAt       *time.Time `gorm:"type:timestamptz"`
	DeliveredAt  *time.Time `gorm:"type:timestamptz"`
	ErrorMessage string     `gorm:"type:text"`
}

func (NotificationLog) TableName() string {
	return "notification_log"
}
