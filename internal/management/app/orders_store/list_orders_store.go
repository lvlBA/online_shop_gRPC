package orders_store

import (
	"context"
	controllersOrdersStore "github.com/lvlBA/online_shop/internal/management/controllers/orders_store"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServiceImpl) ListOrdersStore(ctx context.Context, req *api.ListOrdersStoresRequest) (*api.ListOrdersStoresResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	ordersStore, err := s.ctrlOrdersStore.ListOrderStores(ctx, &controllersOrdersStore.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List orders_stores", err, "request", req)
		return nil, status.Error(codes.Internal, "error list orders_stores")
	}

	result := make([]*api.OrdersStore, 0, len(ordersStore))
	for i := range ordersStore {
		result = append(result, adaptOrdersStoreToApi(ordersStore[i]))
	}
	return &api.ListOrdersStoresResponse{OrdersStore: result}, nil
}
