package user

import (
	"context"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type Service interface {
	CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error)
}

//GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
//ChangePass(ctx context.Context, in *ChangePassRequest, opts ...grpc.CallOption) (*ChangePassResponse, error)
//DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
//	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
