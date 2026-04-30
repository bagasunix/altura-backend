package entities

import "github.com/gofrs/uuid"

type WAAutoReply struct {
	BaseModel
	Name         string
	MatchType    string
	Keywords     []string
	ReplyMessage string
	OnlyNewLead  bool
	SessionIDs   []uuid.UUID
	IsActive     bool
	SortOrder    int
}
