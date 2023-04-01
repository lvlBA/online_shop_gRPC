package cargo

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	"github.com/lvlBA/online_shop/internal/storage/db"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

type CreateParams struct {
	Name         string
	Capacity     int
	Price        float64
	Availability bool
}

func (s *ServiceImpl) CreateCarrier(ctx context.Context, params *CreateParams) (*models.Carrier, error) {
	resp, err := s.db.Cargo().CreateCarrier(ctx, &db.CreateCargoParams{
		Name:         params.Name,
		Capacity:     params.Capacity,
		Price:        params.Price,
		Availability: params.Availability,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
