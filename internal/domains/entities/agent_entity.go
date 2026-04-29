package entities

type Agent struct {
	BaseModel
	Name      string
	Phone     string
	Email     string
	AvatarURL string
	IsActive  bool
	MaxLeads  int
}
