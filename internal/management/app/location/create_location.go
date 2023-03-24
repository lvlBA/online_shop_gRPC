package location

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	controllersLocation "github.com/lvlBA/online_shop/internal/management/controllers/location"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s *ServiceImpl) CreateLocation(ctx context.Context, req *api.CreateLocationRequest) (*api.CreateLocationResponse, error) {
	if err := validateCreateLocationReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	location, err := s.ctrlLocation.CreateLocation(ctx, &controllersLocation.CreateParams{
		Name:     req.Name,
		SiteId:   req.SiteId,
		RegionId: req.RegionId,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "location already exists")
		}
		s.log.Error(ctx, "failed to create location", err, "request", req)

		return nil, status.Error(codes.Internal, "error create location")
	}

	return &api.CreateLocationResponse{
		Location: adaptLocationToApi(location),
	}, nil
}

func validateCreateLocationReq(req *api.CreateLocationRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}

func adaptLocationToApi(model *models.Location) *api.Location {
	return &api.Location{
		Id:   model.ID,
		Name: model.Name,
	}
}
