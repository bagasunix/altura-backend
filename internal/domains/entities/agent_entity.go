package entities

import "time"

type Agent struct {
	ID        string
	Name      string
	Phone     string
	Email     string
	AvatarURL string
	IsActive  bool
	MaxLeads  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
