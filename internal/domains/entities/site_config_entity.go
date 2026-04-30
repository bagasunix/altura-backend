package entities

import (
	"encoding/json"
	"time"
)

type SiteConfig struct {
	ID               int
	WANumber         string
	Email            string
	Tagline          string
	WADefaultMessage string
	HeroCopy         json.RawMessage
	Stats            json.RawMessage
	Instagram        string
	Address          string
	Areas            string
	N8nWebhookURL    string
	AutoReplyEnabled bool
	FollowupEnabled  bool
	AssignAuto       bool
	UpdatedAt        time.Time
}
