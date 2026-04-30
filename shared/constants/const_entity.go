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

type TypeProperty string

const (
	ForSale TypeProperty = "for_sale"
	ForRent TypeProperty = "for_rent"
)

type StatusProperty string

const (
	ActiveProperty    StatusProperty = "active_property"
	NonActiveProperty StatusProperty = "nonactive_property"
)

type ActivityType string

const (
	ActivityTypeStatusChange ActivityType = "status_change"
	ActivityTypeNote         ActivityType = "note"
	ActivityTypeCall         ActivityType = "call"
	ActivityTypeWhatsapp     ActivityType = "whatsapp"
	ActivityTypeEmail        ActivityType = "email"
	ActivityTypeMeeting      ActivityType = "meeting"
	ActivityTypeAutoReply    ActivityType = "auto_reply"
	ActivityTypeAssignment   ActivityType = "assignment"
)

type N8nEventStatus string

const (
	N8nEventStatusPending    N8nEventStatus = "pending"
	N8nEventStatusProcessing N8nEventStatus = "processing"
	N8nEventStatusDone       N8nEventStatus = "done"
	N8nEventStatusFailed     N8nEventStatus = "failed"
)

type NotificationChannel string

const (
	NotificationChannelWhatsapp NotificationChannel = "whatsapp"
	NotificationChannelEmail    NotificationChannel = "email"
	NotificationChannelTelegram NotificationChannel = "telegram"
	NotificationChannelInternal NotificationChannel = "internal"
)

type NotificationStatus string

const (
	NotificationStatusQueued    NotificationStatus = "queued"
	NotificationStatusSent      NotificationStatus = "sent"
	NotificationStatusDelivered NotificationStatus = "delivered"
	NotificationStatusFailed    NotificationStatus = "failed"
)

type WASessionStatus string

const (
	WASessionStatusDisconnected WASessionStatus = "disconnected"
	WASessionStatusConnecting   WASessionStatus = "connecting"
	WASessionStatusConnected    WASessionStatus = "connected"
	WASessionStatusBanned       WASessionStatus = "banned"
	WASessionStatusLoggedOut    WASessionStatus = "logged_out"
)

type WAMessageDirection string

const (
	WAMessageDirectionInbound  WAMessageDirection = "inbound"
	WAMessageDirectionOutbound WAMessageDirection = "outbound"
)

type WAMessageType string

const (
	WAMessageTypeText     WAMessageType = "text"
	WAMessageTypeImage    WAMessageType = "image"
	WAMessageTypeDocument WAMessageType = "document"
	WAMessageTypeAudio    WAMessageType = "audio"
	WAMessageTypeVideo    WAMessageType = "video"
	WAMessageTypeSticker  WAMessageType = "sticker"
	WAMessageTypeLocation WAMessageType = "location"
	WAMessageTypeContact  WAMessageType = "contact"
	WAMessageTypeReaction WAMessageType = "reaction"
	WAMessageTypeTemplate WAMessageType = "template"
	WAMessageTypeUnknown  WAMessageType = "unknown"
)

type WAMessageStatus string

const (
	WAMessageStatusPending   WAMessageStatus = "pending"
	WAMessageStatusSent      WAMessageStatus = "sent"
	WAMessageStatusDelivered WAMessageStatus = "delivered"
	WAMessageStatusRead      WAMessageStatus = "read"
	WAMessageStatusFailed    WAMessageStatus = "failed"
)
