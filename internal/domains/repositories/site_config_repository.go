package repositories

import (
	"context"

	"altura-property/internal/domains/entities"
)

type SiteConfigRepository interface {
	Get(ctx context.Context) (*entities.SiteConfig, error)
	Update(ctx context.Context, cfg *entities.SiteConfig) (*entities.SiteConfig, error)
}
