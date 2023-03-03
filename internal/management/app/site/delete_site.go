package site

import (
	"context"
	"errors"

	"github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) DeleteSite(ctx context.Context, req *api.DeleteSiteRequest) (*api.DeleteSiteResponse, error) {
	if err := validateDeleteSideReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlSite.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, status.Error(codes.Internal, "error delete site")
	}

	return &api.DeleteSiteResponse{}, nil
}

func validateDeleteSideReq(req *api.DeleteSiteRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
