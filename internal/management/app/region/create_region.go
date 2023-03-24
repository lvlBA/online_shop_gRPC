package region

import "C"
import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllersRegion "github.com/lvlBA/online_shop/internal/management/controllers/region"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) CreateRegion(ctx context.Context, req *api.CreateRegionRequest) (*api.CreateRegionResponse, error) {
	if err := validateCreateRegionReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	region, err := s.ctrlRegion.CreateRegion(ctx, &controllersRegion.CreateParams{
		Name:   req.Name,
		SiteId: req.SiteId,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "region already exists")
		}
		s.log.Error(ctx, "failed to create region", err, "request", req)

		return nil, status.Error(codes.Internal, "error create region")
	}

	return &api.CreateRegionResponse{
		Region: adaptRegionToApi(region),
	}, nil
}

func validateCreateRegionReq(req *api.CreateRegionRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}

func adaptRegionToApi(model *models.Region) *api.Region {
	return &api.Region{
		Id:   model.ID,
		Name: model.Name,
	}
}
