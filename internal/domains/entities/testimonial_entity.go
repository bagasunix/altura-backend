package entities

type Testimonial struct {
	BaseModel
	Name      string
	Role      string
	Quote     string
	AvatarURL string
	SortOrder int
	IsActive  bool
}
