package entities

import (
	"time"

	"altura-property/shared/constants"
)

type Lead struct {
	BaseModel
	Name           string
	Phone          string
	City           string
	Type           string
	Budget         string
	Source         string
	Status         constants.LeadStatus
	Priority       constants.LeadPriority
	AssignedTo     *string
	UTMSource      string
	UTMMedium      string
	UTMCampaign    string
	UTMContent     string
	UTMTerm        string
	LastContactAt  *time.Time
	NextFollowupAt *time.Time
	WaSentAt       *time.Time
	QualifiedAt    *time.Time
	ClosedAt       *time.Time
	ListingID      *string
	NotesSummary   string
	IP             string
	UserAgent      string
	Referrer       string
}
