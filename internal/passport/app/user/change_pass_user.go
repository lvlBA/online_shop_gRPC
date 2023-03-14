package user

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s ServiceImpl) ChangePass(ctx context.Context, req *api.ChangePassRequest) (*api.ChangePassResponse, error) {
	if err := validateChangePassUserReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlUser.ChangePass(ctx, req.Id, req.OldPass, req.NewPass); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to change pass", err, "request", req)

		return nil, status.Error(codes.Internal, "error change password")
	}

	return &api.ChangePassResponse{}, nil
}

func validateChangePassUserReq(req *api.ChangePassRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
		"pass": validation.Validate(
			req.OldPass,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^([0-9]{1,}|[a-z]{1,}|[A-Z]{1,}|[!@#$%&*()-_+=]{1,}){8,255}$"),
			),
		),
		"newPass": validation.Validate(
			req.NewPass,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^([0-9]{1,}|[a-z]{1,}|[A-Z]{1,}|[!@#$%&*()-_+=]{1,}){8,255}$"),
			),
		),
	}.Filter()
}
