package auth

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/passport/controllers"
	controllerAuth "github.com/lvlBA/online_shop/internal/passport/controllers/auth"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) CheckUserAccess(ctx context.Context, req *api.CheckUserAccessRequest) (*api.CheckUserAccessResponse, error) {
	if err := validateCheckUserReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token := make([]byte, 0)
	if tokenI := ctx.Value("token"); tokenI != nil {
		t, ok := tokenI.(string)
		if !ok {
			return nil, status.Error(codes.Internal, fmt.Sprintf("failed to type assertion, token type is %s not string", reflect.TypeOf(tokenI).String()))
		}
		token = []byte(t)
	}

	auth, err := s.ctrlAuth.GetUserToken(ctx, &controllerAuth.GetUserTokenRequest{
		Token: token,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.Unauthenticated, "not auth")
		}
		s.log.Error(ctx, "error check user: failed to get user token", err, "request", req)

		return nil, status.Error(codes.Internal, "error check user")
	}

	resource, err := s.ctrlResource.GetResourceByID(ctx, req.ResourceId)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}
		s.log.Error(ctx, "failed to get resource", err, "request", req)

		return nil, status.Error(codes.Internal, "error check user")
	}

	ok, err := s.ctrlAuth.CheckUserAccess(ctx, &controllerAuth.CheckUserAccessRequest{
		UserID:     auth.UserID,
		ResourceID: resource.ID,
	})
	if err != nil {
		s.log.Error(ctx, "failed to check user access", err, "request", req)
		return nil, status.Error(codes.Internal, "error check user")
	}
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "not auth2")
	}

	if err = s.ctrlAuth.UpdateAuth(ctx, auth); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}
		s.log.Error(ctx, "failed to refresh token", err, "request", req)

		return nil, status.Error(codes.Internal, "error check user")

	}

	return &api.CheckUserAccessResponse{}, nil
}

func validateCheckUserReq(req *api.CheckUserAccessRequest) error {
	return validation.Errors{
		"resource_id": validation.Validate(
			req.ResourceId,
			validation.Required,
			is.UUIDv4,
		),
	}.Filter()
}
