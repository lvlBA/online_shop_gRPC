package resource

import (
	"context"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type ServiceImpl struct {
}

func (s *ServiceImpl) CreateResource(ctx context.Context, in *api.CreateResourceRequest) (*api.CreateResourceResponse, error) {
}
func (s *ServiceImpl) GetResource(ctx context.Context, in *api.GetResourceRequest) (*api.GetResourceResponse, error) {
}
func (s *ServiceImpl) DeleteResource(ctx context.Context, in *api.DeleteResourceRequest) (*api.DeleteResourceResponse, error) {
}
func (s *ServiceImpl) ListResource(ctx context.Context, in *api.ListResourceRequest) (*api.ListResourceResponse, error) {
}
func (s *ServiceImpl) SetUserAccess(ctx context.Context, in *api.SetUserAccessRequest) (*api.SetUserAccessResponse, error) {
}
func (s *ServiceImpl) DeleteUserAccess(ctx context.Context, in *api.DeleteUserAccessRequest) (*api.DeleteUserAccessResponse, error) {
}
