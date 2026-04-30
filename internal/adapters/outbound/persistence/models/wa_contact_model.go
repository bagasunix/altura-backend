package models

type WAContact struct {
	BaseModel
	SessionID    string    `gorm:"type:uuid;not null;index"`
	Session      WASession
	JID          string    `gorm:"type:text;not null"`
	Phone        string    `gorm:"type:text;not null;default:''"`
	PushName     string    `gorm:"type:text;not null;default:''"`
	BusinessName string    `gorm:"type:text;not null;default:''"`
	IsBusiness   bool      `gorm:"not null;default:false"`
	IsVerified   bool      `gorm:"not null;default:false"`
	AvatarURL    string    `gorm:"type:text"`
	LeadID       *string   `gorm:"type:uuid;index"`
	Lead         *Lead
}

func (WAContact) TableName() string {
	return "wa_contacts"
}
