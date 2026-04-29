package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type BaseModel struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
