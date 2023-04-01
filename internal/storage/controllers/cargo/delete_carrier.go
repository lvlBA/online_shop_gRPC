package cargo

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
)

func (s *ServiceImpl) DeleteCarrier(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Cargo().DeleteCarrier(ctx, id))
}
