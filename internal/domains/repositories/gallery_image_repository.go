package repositories

import (
	"context"

	"github.com/gofrs/uuid"

	"altura-property/internal/domains/entities"
)

type GalleryImageRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entities.GalleryImage, error)
	FindAll(ctx context.Context, activeOnly bool) ([]*entities.GalleryImage, error)
	Create(ctx context.Context, img *entities.GalleryImage) (*entities.GalleryImage, error)
	Update(ctx context.Context, img *entities.GalleryImage) (*entities.GalleryImage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
