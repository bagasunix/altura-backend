package entities

import "github.com/gofrs/uuid"

type WAContact struct {
	BaseModel
	SessionID    uuid.UUID
	JID          string
	Phone        string
	PushName     string
	BusinessName string
	IsBusiness   bool
	IsVerified   bool
	AvatarURL    string
	LeadID       *uuid.UUID
}
