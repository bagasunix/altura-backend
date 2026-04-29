package models

type Agent struct {
	BaseModel
	Name      string `gorm:"type:text;not null"`
	Phone     string `gorm:"type:text;not null;default:''"`
	Email     string `gorm:"type:text"`
	AvatarURL string `gorm:"type:text;not null;default:''"`
	IsActive  bool   `gorm:"not null;default:true;index"`
	MaxLeads  int    `gorm:"not null;default:30"`
}
