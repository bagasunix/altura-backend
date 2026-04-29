package constants

type LeadStatus string

const (
	LeadStatusNew         LeadStatus = "new"
	LeadStatusContacted   LeadStatus = "contacted"
	LeadStatusQualified   LeadStatus = "qualified"
	LeadStatusSurvey      LeadStatus = "survey"
	LeadStatusNegotiating LeadStatus = "negotiating"
	LeadStatusClosing     LeadStatus = "closing"
	LeadStatusClosedWon   LeadStatus = "closed_won"
	LeadStatusClosedLost  LeadStatus = "closed_lost"
	LeadStatusNurturing   LeadStatus = "nurturing"
)

type LeadPriority string

const (
	LeadPriorityLow    LeadPriority = "low"
	LeadPriorityMedium LeadPriority = "medium"
	LeadPriorityHigh   LeadPriority = "high"
	LeadPriorityUrgent LeadPriority = "urgent"
)
