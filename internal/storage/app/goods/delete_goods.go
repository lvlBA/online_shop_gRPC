package goods

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) DeleteGoods(ctx context.Context, req *api.DeleteGoodsRequest) (*api.DeleteGoodsResponse, error) {
	if err := validateDeleteGoodsReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlGoods.DeleteGoods(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found goods")
		}
		s.log.Error(ctx, "failed to delete goods", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete goods")
	}

	return &api.DeleteGoodsResponse{}, nil
}

func validateDeleteGoodsReq(req *api.DeleteGoodsRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
