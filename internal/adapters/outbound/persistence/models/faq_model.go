package models

type FAQ struct {
	BaseModel
	Question  string `gorm:"type:text;not null"`
	Answer    string `gorm:"type:text;not null"`
	SortOrder int    `gorm:"not null;default:0;index"`
	IsActive  bool   `gorm:"not null;default:true;index"`
}

func (FAQ) TableName() string {
	return "faqs"
}
