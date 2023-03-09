package user

import (
	"context"
	"errors"

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

	site, err := s.ctrlUser.CreateUser(ctx, &controllersUser.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		Sex:       models.Sex(req.Sex),
		Login:     req.Login,
		Password:  req.Pass,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "site already exists")
		}
		s.log.Error(ctx, "failed to create site", err, "request", req)

		return nil, status.Error(codes.Internal, "error create site")
	}

	return &api.CreateUserResponse{
		User: adaptSiteToApi(site),
	}, nil
}

func validateCreateUserReq(req *api.CreateUserRequest) error {
	return validation.Errors{
		"first_name": validation.Validate(req.FirstName, validation.Required), // тут только буквы
		"last_name":  validation.Validate(req.LastName, validation.Required),  // тут только буквы
		"age":        validation.Validate(req.Age, validation.Required),       // от 0 до 200
		"sex":        validation.Validate(req.Sex, validation.Required),       // используем только заданные значения
		"login":      validation.Validate(req.Login, validation.Required),     // ограниччение длинны
		"pass":       validation.Validate(req.Pass, validation.Required),      // сложность пароля
	}.Filter()
}

func adaptSiteToApi(model *models.User) *api.User {
	return &api.User{
		Id:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Age:       model.Age,
		Sex:       api.Sex(model.Sex),
		Login:     model.Login,
	}
}
