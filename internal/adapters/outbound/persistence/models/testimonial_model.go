package models

type Testimonial struct {
	BaseModel
	Name      string `gorm:"type:text;not null"`
	Role      string `gorm:"type:text;not null;default:''"`
	Quote     string `gorm:"type:text;not null"`
	AvatarURL string `gorm:"type:text;not null;default:''"`
	SortOrder int    `gorm:"not null;default:0;index"`
	IsActive  bool   `gorm:"not null;default:true;index"`
}

func (Testimonial) TableName() string {
	return "testimonials"
}
