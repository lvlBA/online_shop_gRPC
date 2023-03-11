package resource

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllersresource "github.com/lvlBA/online_shop/internal/passport/controllers/resource"
	"github.com/lvlBA/online_shop/internal/passport/models"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) CreateResource(ctx context.Context,
	req *api.CreateResourceRequest) (*api.CreateResourceResponse, error) {
	resource, err := s.ctrlService.CreateResource(ctx, &controllersresource.CreateServiceParams{
		Urn: req.Urn,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "Resource is already exists")
		}
		s.log.Error(ctx, "failed to create Resource", err, "request", req)

		return nil, status.Error(codes.Internal, "error create Resource")
	}

	return &api.CreateResourceResponse{
		Resource: adaptResourceToApi(resource),
	}, nil
}

func adaptResourceToApi(model *models.Resource) *api.Resource {
	return &api.Resource{
		Id:  model.ID,
		Urn: model.Urn,
	}
}
