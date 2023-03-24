package orders_store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v7"
	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	v1 "github.com/lvlBA/online_shop/pkg/api/v1"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
	managementclient "github.com/lvlBA/online_shop/pkg/management_client"
	passportapi "github.com/lvlBA/online_shop/pkg/passport/v1"
	passportclient "github.com/lvlBA/online_shop/pkg/passport_client"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type ordersStoreResponse interface {
	GetOrdersStore() *api.OrdersStore
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareOrdersStoreResp(want, got ordersStoreResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareOrdersStore(want.GetOrdersStore(), got.GetOrdersStore())
}

func compareOrdersStore(want, got *api.OrdersStore) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListOrdersStore(want, got *api.ListOrdersStoresResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.OrdersStore) != len(got.OrdersStore) {
		return false
	}
	for i := 0; i < len(want.OrdersStore); i++ {
		if !compareOrdersStore(want.OrdersStore[i], got.OrdersStore[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.OrdersStore) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareOrdersStore(exists[y], got[i]) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func Test_OrdersStore(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("failed to parse config: %s", err)
	}

	intrUserMeta := grpcinterceptors.NewSetUserMeta()

	mCli, err := managementclient.New(ctx, &managementclient.Config{
		Addr: cfg.Addr,
	}, intrUserMeta.GrpcInterceptor)
	if err != nil {
		t.Fatalf("failed to create management: %s", err)
	}

	pCli, err := passportclient.New(ctx, &passportclient.Config{
		Addr: cfg.PassportAddr,
	}, intrUserMeta.GrpcInterceptor)
	if err != nil {
		t.Fatalf("failed to create passport: %s", err)
	}
	user, err := pCli.CreateUser(ctx, &passportapi.CreateUserRequest{
		FirstName: "Andrey",
		LastName:  "Kurochkin",
		Age:       34,
		Sex:       1,
		Login:     "adfhajfasd",
		Pass:      "AAAa1@#a98fuaf",
	})
	if err != nil {
		t.Fatalf("failed to create user: %s", err)
	}
	ctx = context.WithValue(ctx, grpcinterceptors.XRequestIdKey, user.User.Login)
	defer func() {
		if _, err := pCli.DeleteUser(ctx, &passportapi.DeleteUserRequest{
			Id: user.User.Id,
		}); err != nil {
			t.Fatalf("failed to delete user: %s", err)
		}
	}()

	token, err := pCli.GetUserToken(ctx, &passportapi.GetUserTokenRequest{
		Login:    user.User.Login,
		Password: "AAAa1@#a98fuaf",
	})
	if err != nil {
		t.Fatalf("failed to get token: %s", err)
	}
	defer func() {
		// FIXME: удалятьвсе все должно само при удалении пользователя
		if _, err := pCli.DeleteUserToken(ctx, &passportapi.DeleteUserTokenRequest{
			Login:    user.User.Login,
			Password: "AAAa1@#a98fuaf",
		}); err != nil {
			t.Fatalf("faield to delete token: %s", err)
		}
	}()
	ctx = context.WithValue(ctx, grpcinterceptors.TokenKey, token.Token)

	resourceCreateOrdersStore, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.OrdersStoreService/CreateOrdersStore",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateOrdersStore': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateOrdersStore.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateOrdersStore': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateOrdersStore.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateOrdersStore.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteOrdersStore, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.OrdersStoreService/DeleteOrdersStore",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteOrdersStore': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteOrdersStore.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteOrdersStore': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteOrdersStore.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteOrdersStore.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceGetOrdersStore, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.OrdersStoreService/GetOrdersStore",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetOrdersStore': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceGetOrdersStore.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'GetOrdersStore': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceGetOrdersStore.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceGetOrdersStore.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListOrdersStore, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.OrdersStoreService/ListOrdersStore",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListOrdersStore': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListOrdersStore.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListOrdersStore': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListOrdersStore.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListOrdersStore.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceCreateSite, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.SiteService/CreateSite",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateSite': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateSite.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateSite': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateSite.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateSite.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteSite, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.SiteService/DeleteSite",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteSite': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteSite.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteSite': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteSite.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteSite.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	site, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
		Name: "siteName",
	})
	if err != nil {
		t.Fatalf("failed to create site: %s", err)
	}
	defer func() {
		if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: site.Site.Id}); err != nil {
			t.Fatalf("failed to delete site: %s", err)
		}
	}()

	resourceCreateRegion, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.RegionService/CreateRegion",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateRegion': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateRegion.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateRegion': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateRegion.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateRegion.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteRegion, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.RegionService/DeleteRegion",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteRegion': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteRegion.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteRegion': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteRegion.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteRegion.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	region, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
		Name:   "regionName",
		SiteId: site.Site.Id,
	})
	if err != nil {
		t.Fatalf("failed to create region: %s", err)
	}
	defer func() {
		if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: region.Region.Id}); err != nil {
			t.Fatalf("failed to delete region: %s", err)
		}
	}()

	resourceCreateLocation, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.LocationService/CreateLocation",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateLocation': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateLocation.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateLocation': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateLocation.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateLocation.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteLocation, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.LocationService/DeleteLocation",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteLocation': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteLocation.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteLocation': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteLocation.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteLocation.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	location, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
		Name:     "locationName",
		SiteId:   site.Site.Id,
		RegionId: region.Region.Id,
	})
	if err != nil {
		t.Fatalf("failed to create location: %s", err)
	}
	defer func() {
		if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: location.Location.Id}); err != nil {
			t.Fatalf("failed to delete location: %s", err)
		}
	}()

	resourceCreateWarehouse, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.WarehouseService/CreateWarehouse",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateWarehouse': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateWarehouse.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateWarehouse': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateWarehouse.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateWarehouse.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteWarehouse, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.WarehouseService/DeleteWarehouse",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteWarehouse': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteWarehouse.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteWarehouse': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteWarehouse.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteWarehouse.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	warehouse, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
		Name:       "warehouseName",
		SiteId:     site.Site.Id,
		RegionId:   region.Region.Id,
		LocationId: location.Location.Id,
	})
	if err != nil {
		t.Fatalf("failed to create warehouse: %s", err)
	}
	defer func() {
		if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: warehouse.Warehouse.Id}); err != nil {
			t.Fatalf("failed to delete warehouse: %s", err)
		}
	}()

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			ordersStoreName := "test_ordersStore1"
			want := &api.CreateOrdersStoreResponse{OrdersStore: &api.OrdersStore{
				Id:   "",
				Name: ordersStoreName,
			}}

			// check
			got, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
				Name:        ordersStoreName,
				SiteId:      site.Site.Id,
				RegionId:    region.Region.Id,
				LocationId:  location.Location.Id,
				WarehouseId: warehouse.Warehouse.Id,
			})
			if err != nil {
				t.Fatalf("failed to create ordersStore: %s", err)
			}
			want.OrdersStore.Id = got.OrdersStore.Id
			defer func() {
				if _, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: got.OrdersStore.Id}); err != nil {
					t.Fatalf("failed to delete ordersStore: %s", err)
				}
			}()

			if !compareOrdersStoreResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("ordersStore exists", func(t *testing.T) {
				ordersStoreName := "test_ordersStore"
				resp, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
					Name:        ordersStoreName,
					SiteId:      site.Site.Id,
					RegionId:    region.Region.Id,
					LocationId:  location.Location.Id,
					WarehouseId: warehouse.Warehouse.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: resp.OrdersStore.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
					Name:        ordersStoreName,
					SiteId:      site.Site.Id,
					RegionId:    region.Region.Id,
					LocationId:  location.Location.Id,
					WarehouseId: warehouse.Warehouse.Id,
				})
				if err == nil {
					defer func() {
						if _, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: resp2.OrdersStore.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double ordersStore creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
							Name: "",
						})

						assert.Equalf(t, status.Code(err), codes.InvalidArgument,
							"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
						assert.Nil(t, resp)
					})
				})
			})
		})
	})
	t.Run("get_Success", func(t *testing.T) {
		// define
		ordersStore, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
			Name:        "test",
			SiteId:      site.Site.Id,
			RegionId:    region.Region.Id,
			LocationId:  location.Location.Id,
			WarehouseId: warehouse.Warehouse.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetOrdersStoreResponse{OrdersStore: &api.OrdersStore{
			Id:   ordersStore.OrdersStore.Id,
			Name: "test",
		}}

		// check
		got, err := mCli.GetOrdersStore(ctx, &api.GetOrdersStoreRequest{Id: ordersStore.OrdersStore.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareOrdersStoreResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: got.OrdersStore.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("ordersStore_doesn't exist", func(t *testing.T) {
			t.Run("ordersStore_is_bad", func(t *testing.T) {
				resp, err := mCli.GetOrdersStore(ctx, &api.GetOrdersStoreRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("orders_store_is_empty", func(t *testing.T) {
				resp, err := mCli.GetOrdersStore(ctx, &api.GetOrdersStoreRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		ordersStore, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
			Name:        "test",
			SiteId:      site.Site.Id,
			RegionId:    region.Region.Id,
			LocationId:  location.Location.Id,
			WarehouseId: warehouse.Warehouse.Id,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: ordersStore.OrdersStore.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("orders_store_doesn't exist", func(t *testing.T) {
			t.Run("orders_store_is_bad", func(t *testing.T) {
				resp, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("orders_store_is_empty", func(t *testing.T) {
				resp, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("list_Orders_store_pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define
				ordersStore, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
					Name:        "test",
					SiteId:      site.Site.Id,
					RegionId:    region.Region.Id,
					LocationId:  location.Location.Id,
					WarehouseId: warehouse.Warehouse.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListOrdersStoresResponse{OrdersStore: []*api.OrdersStore{
					{
						Id:   ordersStore.OrdersStore.Id,
						Name: "test",
					},
				}}

				// check
				got, err := mCli.ListOrdersStore(ctx, &api.ListOrdersStoresRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListOrdersStore(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: ordersStore.OrdersStore.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.OrdersStore, 0, 20)
				for i := 0; i < len(exists); i++ {
					ordersStoreName := fmt.Sprint(i + 252)
					got, err := mCli.CreateOrdersStore(ctx, &api.CreateOrdersStoreRequest{
						Name:        ordersStoreName,
						SiteId:      site.Site.Id,
						RegionId:    region.Region.Id,
						LocationId:  location.Location.Id,
						WarehouseId: warehouse.Warehouse.Id,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.OrdersStore)
					defer func() {
						if _, err := mCli.DeleteOrdersStore(ctx, &api.DeleteOrdersStoreRequest{Id: got.OrdersStore.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := mCli.ListOrdersStore(ctx, &api.ListOrdersStoresRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.OrdersStore) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.OrdersStore) < len(exists)")
						case limit < len(exists) && len(got.OrdersStore) != limit:
							t.Fatalf(" limit < len(exists) && len(got.OrdersStore) != limit")
						case !notLineCompare(exists, got.OrdersStore):
							t.Fatalf("!notLineCompare(exists, got.OrdersStore)")
						}
					})
				}
			})
			t.Run("case_page_two_limit_one", func(t *testing.T) {
				// TODO: tests
			})
		})

	})

}
