package models

type LeadNote struct {
	BaseModel
	LeadID string `gorm:"type:uuid;not null;index"`
	Lead   Lead
	Body   string `gorm:"type:text;not null"`
}

func (LeadNote) TableName() string {
	return "lead_notes"
}
