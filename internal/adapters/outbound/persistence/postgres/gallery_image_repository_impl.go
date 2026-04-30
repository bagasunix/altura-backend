package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type galleryImageRepository struct {
	db *gorm.DB
}

func NewGalleryImageRepository(db *gorm.DB) repositories.GalleryImageRepository {
	return &galleryImageRepository{db: db}
}

func (r *galleryImageRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.GalleryImage, error) {
	panic("unimplemented")
}

func (r *galleryImageRepository) FindAll(ctx context.Context, activeOnly bool) ([]*entities.GalleryImage, error) {
	panic("unimplemented")
}

func (r *galleryImageRepository) Create(ctx context.Context, img *entities.GalleryImage) (*entities.GalleryImage, error) {
	panic("unimplemented")
}

func (r *galleryImageRepository) Update(ctx context.Context, img *entities.GalleryImage) (*entities.GalleryImage, error) {
	panic("unimplemented")
}

func (r *galleryImageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
