package entities

type FAQ struct {
	BaseModel
	Question  string
	Answer    string
	SortOrder int
	IsActive  bool
}
