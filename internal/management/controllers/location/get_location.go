package location

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/models"
)

func (s *ServiceImpl) GetLocation(ctx context.Context, id string) (*models.Location, error) {
	resp, err := s.db.Location().GetLocation(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
