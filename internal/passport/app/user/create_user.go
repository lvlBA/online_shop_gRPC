package user

import (
	"context"
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllersUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	"github.com/lvlBA/online_shop/internal/passport/models"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s ServiceImpl) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	if err := validateCreateUserReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.ctrlUser.CreateUser(ctx, &controllersUser.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Sex:       models.Sex(req.Sex),
		Login:     req.Login,
		Password:  req.Pass,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		s.log.Error(ctx, "failed to create user", err, "request", req)

		return nil, status.Error(codes.Internal, "error create user")
	}

	return &api.CreateUserResponse{

		User: adaptUserToApi(user),
	}, nil
}

func validateCreateUserReq(req *api.CreateUserRequest) error {
	return validation.Errors{
		"first_name": validation.Validate(
			req.FirstName,
			validation.Required,
			validation.Match(regexp.MustCompile("^[a-zA-Z]{3,255}$")),
		),
		"last_name": validation.Validate(
			req.LastName,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^[a-zA-Z]{3,255}$"),
			),
		),
		"age": validation.Validate(
			int64(req.Age),
			validation.Required,
			validation.Min(1),
			validation.Max(200),
		),
		"sex": validation.Validate(
			req.Sex.String(),
			validation.Required,
			validation.In(api.Sex_SexFemale.String(), api.Sex_SexMale.String()),
		),
		"login": validation.Validate(
			req.Login,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]{4,254}$"),
			),
		),
		"pass": validation.Validate(
			req.Pass,
			validation.Required,
			validation.Match(
				regexp.MustCompile("^([0-9]{1,}|[a-z]{1,}|[A-Z]{1,}|[!@#$%&*()-_+=]{1,}){8,255}$"),
			),
		),
	}.Filter()
}

func adaptUserToApi(model *models.User) *api.User {
	return &api.User{
		Id:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Age:       model.Age,
		Sex:       api.Sex(model.Sex),
		Login:     model.Login,
	}
}
