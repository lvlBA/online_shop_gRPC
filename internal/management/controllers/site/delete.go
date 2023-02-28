package site

import "context"

func (s *ServiceImpl) Delete(ctx context.Context, id string) error {
	return s.db.Site().DeleteSite(ctx, id)

}
