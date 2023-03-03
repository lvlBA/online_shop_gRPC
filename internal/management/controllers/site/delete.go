package site

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) Delete(ctx context.Context, id string) (err error) {
	return controllers.AdaptingErrorDB(s.db.Site().DeleteSite(ctx, id))

}
