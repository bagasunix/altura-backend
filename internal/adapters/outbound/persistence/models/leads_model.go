package models

import "time"

type Lead struct {
	BaseModel

	Name           string         `gorm:"type:text;not null"`
	Phone          string         `gorm:"type:text;not null"`
	City           string         `gorm:"type:text"`
	Type           string         `gorm:"type:text"`
	Budget         string         `gorm:"type:text"`
	Source         string         `gorm:"type:text"`
	Status         string         `gorm:"type:lead_status;not null;default:'new';index"`
	Priority       string         `gorm:"type:lead_priority;not null;default:'medium';index"`
	AssignedTo     *string        `gorm:"type:uuid;index"`
	Agent          *Agent         `gorm:"foreignKey:AssignedTo"`
	UTMSource      string         `gorm:"type:text;index"`
	UTMMedium      string         `gorm:"type:text"`
	UTMCampaign    string         `gorm:"type:text"`
	UTMContent     string         `gorm:"type:text"`
	UTMTerm        string         `gorm:"type:text"`
	LastContactAt  *time.Time     `gorm:"type:timestamptz"`
	NextFollowupAt *time.Time     `gorm:"type:timestamptz;index"`
	WaSentAt       *time.Time     `gorm:"type:timestamptz"`
	QualifiedAt    *time.Time     `gorm:"type:timestamptz"`
	ClosedAt       *time.Time     `gorm:"type:timestamptz"`
	ListingID      *string        `gorm:"type:uuid"`
	Listing        *PropertyModel `gorm:"foreignKey:ListingID"`
	NotesSummary   string         `gorm:"type:text"`
	IP             string         `gorm:"type:text"`
	UserAgent      string         `gorm:"type:text"`
	Referrer       string         `gorm:"type:text"`
}
