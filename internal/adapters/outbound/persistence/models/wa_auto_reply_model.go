package models

type WAAutoReply struct {
	BaseModel
	Name         string   `gorm:"type:text;not null"`
	MatchType    string   `gorm:"type:text;not null;default:'contains'"`
	Keywords     []string `gorm:"type:text[]"`
	ReplyMessage string   `gorm:"type:text;not null"`
	OnlyNewLead  bool     `gorm:"not null;default:true"`
	SessionIDs   []string `gorm:"type:uuid[]"`
	IsActive     bool     `gorm:"not null;default:true"`
	SortOrder    int      `gorm:"not null;default:0;index"`
}

func (WAAutoReply) TableName() string {
	return "wa_auto_replies"
}
