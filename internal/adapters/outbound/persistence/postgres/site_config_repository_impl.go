package postgres

import (
	"context"

	"gorm.io/gorm"

	"altura-property/internal/domains/entities"
	"altura-property/internal/domains/repositories"
)

type siteConfigRepository struct {
	db *gorm.DB
}

func NewSiteConfigRepository(db *gorm.DB) repositories.SiteConfigRepository {
	return &siteConfigRepository{db: db}
}

func (r *siteConfigRepository) Get(ctx context.Context) (*entities.SiteConfig, error) {
	panic("unimplemented")
}

func (r *siteConfigRepository) Update(ctx context.Context, cfg *entities.SiteConfig) (*entities.SiteConfig, error) {
	panic("unimplemented")
}
