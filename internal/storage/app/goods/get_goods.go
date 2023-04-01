package goods

import (
	"context"
	"errors"
	controllerGoods "github.com/lvlBA/online_shop/internal/storage/controllers/goods"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) GetGoods(ctx context.Context, req *api.GetGoodsRequest) (*api.GetGoodsResponse, error) {
	if err := validateGetGoodsReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	goods, err := s.ctrlGoods.GetGoods(ctx, &controllerGoods.GetGoodsParams{
		ID:   &req.Id,
		Name: &req.Name,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "goods not found")
		}
		s.log.Error(ctx, "failed to get goods", err, "request", req)

		return nil, status.Error(codes.Internal, "error get goods")
	}

	return &api.GetGoodsResponse{
		Goods: adaptGoodsToApi(goods),
	}, nil

}

func validateGetGoodsReq(req *api.GetGoodsRequest) error {
	return validation.Errors{
		"id":   validation.Validate(req.Id, validation.Required, is.UUIDv4),
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}
