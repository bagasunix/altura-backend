package entities

import "github.com/shopspring/decimal"

type WATemplate struct {
	BaseModel
	Name         string
	TriggerEvent string
	Message      string
	DelayHours   decimal.Decimal
	IsActive     bool
	SortOrder    int
}
