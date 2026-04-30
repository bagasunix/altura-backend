package models

type GalleryImage struct {
	BaseModel
	Src       string `gorm:"type:text;not null"`
	Alt       string `gorm:"type:text;not null;default:''"`
	SortOrder int    `gorm:"not null;default:0;index"`
	IsActive  bool   `gorm:"not null;default:true;index"`
}

func (GalleryImage) TableName() string {
	return "gallery_images"
}
