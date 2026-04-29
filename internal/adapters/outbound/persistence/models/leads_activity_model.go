package models

type LeadsActivity struct {
	BaseModel
	LeadID   string `gorm:"type:uuid;not null;index"`
	Lead     Lead
	AgentID  *string `gorm:"type:uuid"`
	Agent    *Agent
	Type     string                 `gorm:"type:activity_type;not null;default:'note';index"`
	Body     string                 `gorm:"type:text;not null;default:''"`
	Metadata map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
	IsAuto   bool                   `gorm:"not null;default:false;index"`
}
