package orders_store

import (
	"context"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllersOrdersStore "github.com/lvlBA/online_shop/internal/management/controllers/orders_store"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServiceImpl) CreateOrdersStore(ctx context.Context, req *api.CreateOrdersStoreRequest) (*api.CreateOrdersStoreResponse, error) {

	if err := validateCreateOrdersStoreReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ordersStore, err := s.ctrlOrdersStore.CreateOrderStore(ctx, &controllersOrdersStore.CreateParams{
		Name:        req.Name,
		SiteId:      req.SiteId,
		RegionId:    req.RegionId,
		LocationId:  req.LocationId,
		WarehouseId: req.WarehouseId,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "orders_store already exists")
		}
		s.log.Error(ctx, "failed to create orders_store", err, "request", req)

		return nil, status.Error(codes.Internal, "error create orders_store")
	}

	return &api.CreateOrdersStoreResponse{
		OrdersStore: adaptOrdersStoreToApi(ordersStore),
	}, nil
}

func validateCreateOrdersStoreReq(req *api.CreateOrdersStoreRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}

func adaptOrdersStoreToApi(model *models.OrdersStore) *api.OrdersStore {
	return &api.OrdersStore{
		Id:   model.ID,
		Name: model.Name,
	}
}
