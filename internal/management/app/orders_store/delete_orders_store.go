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

func (s *ServiceImpl) DeleteOrdersStore(ctx context.Context, req *api.DeleteOrdersStoreRequest) (*api.DeleteOrdersStoreResponse, error) {
	if err := validateDeleteOrdersStoreReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlOrdersStore.DeleteOrderStore(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete orders_store", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete orders_store")
	}

	return &api.DeleteOrdersStoreResponse{}, nil
}

func validateDeleteOrdersStoreReq(req *api.DeleteOrdersStoreRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
