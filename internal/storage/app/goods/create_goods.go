package goods

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	controllersGoods "github.com/lvlBA/online_shop/internal/storage/controllers/goods"
	"github.com/lvlBA/online_shop/internal/storage/models"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) CreateGoods(ctx context.Context, req *api.CreateGoodsRequest) (*api.CreateGoodsResponse, error) {
	if err := validateCreateGoodsReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	goods, err := s.ctrlGoods.CreateGoods(ctx, &controllersGoods.CreateParams{
		Name:   req.Name,
		Weight: int(req.Weight),
		Length: int(req.Length),
		Width:  int(req.Width),
		Height: int(req.Height),
		Price:  float64(req.Price),
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "goods already exists")
		}
		s.log.Error(ctx, "failed to create goods", err, "request", req)

		return nil, status.Error(codes.Internal, "error create goods")
	}

	return &api.CreateGoodsResponse{
		Goods: adaptGoodsToApi(goods),
	}, nil
}

func validateCreateGoodsReq(req *api.CreateGoodsRequest) error {
	return validation.Errors{
		"name":   validation.Validate(req.Name, validation.Required),
		"weight": validation.Validate(req.Weight, validation.Required),
		"length": validation.Validate(req.Length, validation.Required),
		"width":  validation.Validate(req.Width, validation.Required),
		"height": validation.Validate(req.Height, validation.Required),
		"price":  validation.Validate(req.Price, validation.Required),
	}.Filter()
}

func adaptGoodsToApi(model *models.Goods) *api.Goods {
	return &api.Goods{
		Id:     model.ID,
		Name:   model.Name,
		Weight: uint64(model.Weight),
		Length: uint64(model.Length),
		Width:  uint64(model.Width),
		Height: uint64(model.Height),
		Price:  float32(model.Price),
	}
}
