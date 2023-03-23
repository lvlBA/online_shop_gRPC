package location

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteLocation(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Location().DeleteLocation(ctx, id))
}
