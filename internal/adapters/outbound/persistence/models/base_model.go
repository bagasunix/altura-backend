package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create"`
	CreatedAt time.Time `gorm:"autoCreateTime;index;sort:desc;<-:create"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;<-:update"`
}
