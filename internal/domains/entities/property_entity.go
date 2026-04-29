package entities

import "time"

type Property struct {
	ID       string
	Title    string
	Location string
	District string
	Price    string
	Beds     int
	Baths    int
	Area     int
	LandArea int
	Image    string
	Images   []string

	Badge        string
	Type         string
	PropertyType string
	Certificate  string
	Description  string

	IsActive     bool
	IsFeatured   bool
	Views        int
	InquiryCount int

	CreatedAt time.Time
	UpdatedAt time.Time
}
