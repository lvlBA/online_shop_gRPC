package region

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

func (s ServiceImpl) DeleteRegion(ctx context.Context, req *api.DeleteRegionRequest) (*api.DeleteRegionResponse, error) {
	if err := validateDeleteRegionReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlRegion.DeleteRegion(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete region", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete region")
	}

	return &api.DeleteRegionResponse{}, nil
}

func validateDeleteRegionReq(req *api.DeleteRegionRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
