package entities

import "github.com/shopspring/decimal"

type Property struct {
	BaseModel
	Title        string
	Location     string
	District     string
	Price        decimal.Decimal
	Beds         int
	Baths        int
	Area         int
	LandArea     int
	Image        string
	Images       []string
	Badge        string
	Type         string
	PropertyType string
	Certificate  string
	Description  string
	IsActive     bool
	IsFeatured   bool
	Views        int
	InquiryCount int
}
