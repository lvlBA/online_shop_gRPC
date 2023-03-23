package region

import (
	"context"
	"errors"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) GetRegion(ctx context.Context, req *api.GetRegionRequest) (*api.GetRegionResponse, error) {
	if err := validateGetRegionReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	region, err := s.ctrlRegion.GetRegion(ctx, req.Id)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "region not found")
		}
		s.log.Error(ctx, "failed to get region", err, "request", req)

		return nil, status.Error(codes.Internal, "error get region")
	}

	return &api.GetRegionResponse{
		Region: adaptRegionToApi(region),
	}, nil
}

func validateGetRegionReq(req *api.GetRegionRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
