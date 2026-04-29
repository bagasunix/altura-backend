package entities

import (
	"time"

	"altura-property/shared/constant"
)

type Lead struct {
	ID     string
	Name   string
	Phone  string
	City   string
	Type   string
	Budget string
	Source string

	Status     constant.LeadStatus
	Priority   constant.LeadPriority
	AssignedTo *string

	UTMSource   string
	UTMMedium   string
	UTMCampaign string
	UTMContent  string
	UTMTerm     string

	LastContactAt  *time.Time
	NextFollowupAt *time.Time
	WaSentAt       *time.Time
	QualifiedAt    *time.Time
	ClosedAt       *time.Time

	ListingID    *string
	NotesSummary string

	IP        string
	UserAgent string
	Referrer  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
