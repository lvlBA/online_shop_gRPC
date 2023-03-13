package auth

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllerAuth "github.com/lvlBA/online_shop/internal/passport/controllers/auth"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) CheckUser(ctx context.Context, req *api.CheckUserAccessRequest) (*api.CheckUserAccessResponse, error) {
	if err := validateCheckUserReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.ctrlAuth.GetUserToken(ctx, &controllerAuth.GetUserTokenRequest{
		Token: &req.Token,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.Unauthenticated, "not auth")
		}
		s.log.Error(ctx, "failed to get user token", err, "request", req)

		return nil, status.Error(codes.Internal, "error check user")
	}

	resource, err := s.ctrlResource.GetResourceByName(ctx, req.Resource)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}
		s.log.Error(ctx, "failed to get resource", err, "request", req)

		return nil, status.Error(codes.Internal, "error check user")
	}

	ok, err := s.ctrlAuth.CheckUserAccess(ctx, &controllerAuth.CheckUserAccessRequest{
		UserID:     user.ID,
		ResourceID: resource.ID,
	})
	if err != nil {
		s.log.Error(ctx, "failed to check user access", err, "request", req)
		return nil, status.Error(codes.Internal, "error check user")
	}

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "not auth")
	}

	return &api.CheckUserAccessResponse{}, nil
}

func validateCheckUserReq(req *api.CheckUserAccessRequest) error {
	return validation.Errors{
		"resource_id": validation.Validate(
			req.Resource,
			validation.Required,
			is.UUIDv4,
		),
		"token": validation.Validate(
			req.Token,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^([0-9]{1,}|[a-z]{1,}|[A-Z]{1,}|[!@#$%&*()-_+=]{1,}){8,255}$"),
			),
		),
	}.Filter()
}
