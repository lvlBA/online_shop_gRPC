package auth

import (
	"context"
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) DeleteUserToken(ctx context.Context, req *api.DeleteUserTokenRequest) (*api.DeleteUserTokenResponse, error) {
	if err := validateDeleteTokenReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlAuth.DeleteUserToken(ctx, req.Token); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete token", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete token")
	}

	return &api.DeleteUserTokenResponse{}, nil
}

func validateDeleteTokenReq(req *api.DeleteUserTokenRequest) error {
	return validation.Errors{
		"token": validation.Validate(
			req.Token,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^([0-9]{1,}|[a-z]{1,}|[A-Z]{1,}|[!@#$%&*()-_+=]{1,}){8,255}$"),
			),
		),
	}.Filter()
}
