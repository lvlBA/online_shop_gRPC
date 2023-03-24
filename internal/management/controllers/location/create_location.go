package location

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type CreateParams struct {
	Name     string
	SiteId   string
	RegionId string
}

func (s *ServiceImpl) CreateLocation(ctx context.Context, params *CreateParams) (*models.Location, error) {
	resp, err := s.db.Location().CreateLocation(ctx, &db.CreateLocationParams{
		Name:     params.Name,
		SiteId:   params.SiteId,
		RegionId: params.RegionId,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
