package models

import "github.com/shopspring/decimal"

type WATemplate struct {
	BaseModel
	Name         string          `gorm:"type:text;not null;unique"`
	TriggerEvent string          `gorm:"type:text;not null;index"`
	Message      string          `gorm:"type:text;not null"`
	DelayHours   decimal.Decimal `gorm:"type:numeric(5,2);not null;default:0"`
	IsActive     bool            `gorm:"not null;default:true;index"`
	SortOrder    int             `gorm:"not null;default:0"`
}

func (WATemplate) TableName() string {
	return "wa_templates"
}
