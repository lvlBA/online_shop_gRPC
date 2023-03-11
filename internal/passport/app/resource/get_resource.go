package resource

import (
	"context"
	"errors"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServiceImpl) GetResource(ctx context.Context, req *api.GetResourceRequest) (*api.GetResourceResponse, error) {
	if err := validateGetResourceReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resource, err := s.ctrlService.GetResource(ctx, req.Id)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "site not found")
		}
		s.log.Error(ctx, "failed to get Resource", err, "request", req)

		return nil, status.Error(codes.Internal, "error get resource")
	}

	return &api.GetResourceResponse{
		Resource: adaptResourceToApi(resource),
	}, nil
}

func validateGetResourceReq(req *api.GetResourceRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
