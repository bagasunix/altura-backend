package models

import "github.com/shopspring/decimal"

type PropertyModel struct {
	BaseModel
	Title        string          `gorm:"type:text;not null"`
	Location     string          `gorm:"type:text;not null"`
	District     string          `gorm:"type:text;not null"`
	Price        decimal.Decimal `gorm:"type:numeric(15,2);not null"`
	Beds         int             `gorm:"not null;default:0"`
	Baths        int             `gorm:"not null;default:0"`
	Area         int             `gorm:"not null;default:0"`
	LandArea     int             `gorm:"not null;default:0"`
	Image        string          `gorm:"type:text;not null;default:''"`
	Images       []string        `gorm:"type:text[]"`
	Badge        string          `gorm:"type:text"`
	Type         string          `gorm:"type:text;not null"`
	PropertyType string          `gorm:"type:text;not null;default:'rumah'"`
	Certificate  string          `gorm:"type:text;not null;default:''"`
	Description  string          `gorm:"type:text;not null;default:''"`
	IsActive     bool            `gorm:"not null;default:true;index"`
	IsFeatured   bool            `gorm:"not null;default:false;index"`
	Views        int             `gorm:"not null;default:0"`
	InquiryCount int             `gorm:"not null;default:0"`
}

func (PropertyModel) TableName() string {
	return "properties"
}
