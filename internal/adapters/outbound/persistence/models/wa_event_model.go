package models

import "time"

type N8nEvent struct {
	BaseModel
	EventType    string                 `gorm:"type:text;not null;index"`
	EntityType   string                 `gorm:"type:text;not null;index"`
	EntityID     *string                `gorm:"type:uuid"`
	Payload      map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
	Status       string                 `gorm:"type:n8n_event_status;not null;default:'pending';index"`
	Attempts     int                    `gorm:"not null;default:0"`
	MaxAttempts  int                    `gorm:"not null;default:3"`
	ScheduledAt  time.Time              `gorm:"type:timestamptz;not null;default:now();index"`
	ProcessedAt  *time.Time             `gorm:"type:timestamptz"`
	ErrorMessage string                 `gorm:"type:text"`
}
