package goods

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	controllerGoods "github.com/lvlBA/online_shop/internal/storage/controllers/goods"
	"github.com/lvlBA/online_shop/internal/storage/models"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) ListGoods(ctx context.Context, req *api.ListGoodsRequest) (*api.ListGoodsResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	goods, err := s.ctrlGoods.ListGoods(ctx, &controllerGoods.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List goods", err, "request", req)
		return nil, status.Error(codes.Internal, "error list goods")
	}

	result := make([]*api.Goods, 0, len(goods))
	for i := range goods {
		result = append(result, adaptGoodsToApi(goods[i]))
	}
	return &api.ListGoodsResponse{Goods: result}, nil
}
