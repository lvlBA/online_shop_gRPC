package site

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) CreateSite(ctx context.Context, req *api.CreateSideRequest) (*api.CreateSideResponse, error) {
	if err := validateCreateSideReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	site, err := s.ctrlSite.Create(ctx, &controllersSite.CreateParams{
		Name: req.Name,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "site already exists")
		}
		s.log.Error(ctx, "failed to create site", err, "request", req)

		return nil, status.Error(codes.Internal, "error create site")
	}

	return &api.CreateSideResponse{
		Site: adaptSiteToApi(site),
	}, nil
}

func validateCreateSideReq(req *api.CreateSideRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}

func adaptSiteToApi(model *models.Site) *api.Site {
	return &api.Site{
		Id:   model.ID,
		Name: model.Name,
	}
}
