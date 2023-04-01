package cargo

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	"github.com/lvlBA/online_shop/internal/storage/db"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

type GetCargoParams struct {
	ID   *string
	Name *string
}

func (s *ServiceImpl) GetCarrier(ctx context.Context, params *GetCargoParams) (*models.Carrier, error) {
	carrier, err := s.db.Cargo().GetCarrier(ctx, &db.GetCargoParams{
		UserID: params.ID,
		Name:   params.Name,
	})
	if err != nil {
		err = controllers.AdaptingErrorDB(err)
		return nil, err
	}

	return carrier, nil
}
