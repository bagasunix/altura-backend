package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type GalleryImage struct {
	ID        uuid.UUID
	Src       string
	Alt       string
	SortOrder int
	IsActive  bool
	CreatedAt time.Time
}
