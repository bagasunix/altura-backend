package models

import (
	"encoding/json"
	"time"
)

type SiteConfig struct {
	ID               int             `gorm:"primaryKey;default:1"`
	WANumber         string          `gorm:"type:text;not null;default:'628123456789'"`
	Email            string          `gorm:"type:text;not null;default:'halo@altura-prop.id'"`
	Tagline          string          `gorm:"type:text;not null;default:''"`
	WADefaultMessage string          `gorm:"type:text;not null;default:''"`
	HeroCopy         json.RawMessage `gorm:"type:jsonb;not null;default:'{}'"`
	Stats            json.RawMessage `gorm:"type:jsonb;not null;default:'[]'"`
	Instagram        string          `gorm:"type:text;not null;default:''"`
	Address          string          `gorm:"type:text;not null;default:''"`
	Areas            string          `gorm:"type:text;not null;default:''"`
	N8nWebhookURL    string          `gorm:"column:n8n_webhook_url;type:text;not null;default:''"`
	AutoReplyEnabled bool            `gorm:"not null;default:true"`
	FollowupEnabled  bool            `gorm:"not null;default:true"`
	AssignAuto       bool            `gorm:"not null;default:false"`
	UpdatedAt        time.Time       `gorm:"autoUpdateTime"`
}

func (SiteConfig) TableName() string {
	return "site_config"
}
