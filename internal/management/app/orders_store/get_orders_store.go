package orders_store

import (
	"context"
	"errors"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s *ServiceImpl) GetOrdersStore(ctx context.Context, req *api.GetOrdersStoreRequest) (*api.GetOrdersStoreResponse, error) {
	if err := validateGetOrdersStoreReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ordersStore, err := s.ctrlOrdersStore.GetOrderStore(ctx, req.Id)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "orders_store not found")
		}
		s.log.Error(ctx, "failed to get orders_store", err, "request", req)

		return nil, status.Error(codes.Internal, "error get orders_store")
	}

	return &api.GetOrdersStoreResponse{
		OrdersStore: adaptOrdersStoreToApi(ordersStore),
	}, nil
}

func validateGetOrdersStoreReq(req *api.GetOrdersStoreRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
