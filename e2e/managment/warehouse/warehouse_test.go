package warehouse

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

type warehouseResponse interface {
	GetWarehouse() *api.Warehouse
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareWarehouseResp(want, got warehouseResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareWarehouse(want.GetWarehouse(), got.GetWarehouse())
}

func compareWarehouse(want, got *api.Warehouse) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListWarehouses(want, got *api.ListWarehousesResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Warehouse) != len(got.Warehouse) {
		return false
	}
	for i := 0; i < len(want.Warehouse); i++ {
		if !compareWarehouse(want.Warehouse[i], got.Warehouse[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.Warehouse) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareWarehouse(exists[y], got[i]) {
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

func Test_Warehouse(t *testing.T) {
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

	resourceGetWarehouse, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.WarehouseService/GetWarehouse",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetWarehouse': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceGetWarehouse.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'GetWarehouse': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceGetWarehouse.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceGetWarehouse.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListWarehouse, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.WarehouseService/ListWarehouse",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListWarehouse': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListWarehouse.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListWarehouse': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListWarehouse.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListWarehouse.Resource.Id,
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

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			warehouseName := "test_warehouse1"
			want := &api.CreateWarehouseResponse{Warehouse: &api.Warehouse{
				Id:   "",
				Name: warehouseName,
			}}

			// check
			got, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
				Name:       warehouseName,
				SiteId:     site.Site.Id,
				RegionId:   region.Region.Id,
				LocationId: location.Location.Id,
			})
			if err != nil {
				t.Fatalf("failed to create warehouse: %s", err)
			}
			want.Warehouse.Id = got.Warehouse.Id
			defer func() {
				if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: got.Warehouse.Id}); err != nil {
					t.Fatalf("failed to delete warehouse: %s", err)
				}
			}()

			if !compareWarehouseResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("warehouse exists", func(t *testing.T) {
				warehouseName := "test_warehouse"
				resp, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
					Name:       warehouseName,
					SiteId:     site.Site.Id,
					RegionId:   region.Region.Id,
					LocationId: location.Location.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: resp.Warehouse.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
					Name:       warehouseName,
					SiteId:     site.Site.Id,
					RegionId:   region.Region.Id,
					LocationId: location.Location.Id,
				})
				if err == nil {
					defer func() {
						if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: resp2.Warehouse.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double warehouse creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
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
		warehouse, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
			Name:       "test",
			SiteId:     site.Site.Id,
			RegionId:   region.Region.Id,
			LocationId: location.Location.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetWarehouseResponse{Warehouse: &api.Warehouse{
			Id:   warehouse.Warehouse.Id,
			Name: "test",
		}}

		// check
		got, err := mCli.GetWarehouse(ctx, &api.GetWarehouseRequest{Id: warehouse.Warehouse.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareWarehouseResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: got.Warehouse.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("warehouse_doesn't exist", func(t *testing.T) {
			t.Run("warehouse_is_bad", func(t *testing.T) {
				resp, err := mCli.GetWarehouse(ctx, &api.GetWarehouseRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("warehouse_is_empty", func(t *testing.T) {
				resp, err := mCli.GetWarehouse(ctx, &api.GetWarehouseRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		warehouse, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
			Name:       "test",
			SiteId:     site.Site.Id,
			RegionId:   region.Region.Id,
			LocationId: location.Location.Id,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: warehouse.Warehouse.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("warehouse_doesn't exist", func(t *testing.T) {
			t.Run("warehouse_is_bad", func(t *testing.T) {
				resp, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("warehouse_is_empty", func(t *testing.T) {
				resp, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("listWarehouse_pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define
				warehouse, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
					Name:       "test",
					SiteId:     site.Site.Id,
					RegionId:   region.Region.Id,
					LocationId: location.Location.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListWarehousesResponse{Warehouse: []*api.Warehouse{
					{
						Id:   warehouse.Warehouse.Id,
						Name: "test",
					},
				}}

				// check
				got, err := mCli.ListWarehouse(ctx, &api.ListWarehousesRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListWarehouses(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: warehouse.Warehouse.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.Warehouse, 0, 20)
				for i := 0; i < len(exists); i++ {
					warehouseName := fmt.Sprint(i + 252)
					got, err := mCli.CreateWarehouse(ctx, &api.CreateWarehouseRequest{
						Name:       warehouseName,
						SiteId:     site.Site.Id,
						RegionId:   region.Region.Id,
						LocationId: location.Location.Id,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Warehouse)
					defer func() {
						if _, err := mCli.DeleteWarehouse(ctx, &api.DeleteWarehouseRequest{Id: got.Warehouse.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := mCli.ListWarehouse(ctx, &api.ListWarehousesRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Warehouse) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Warehouse) < len(exists)")
						case limit < len(exists) && len(got.Warehouse) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Warehouse) != limit")
						case !notLineCompare(exists, got.Warehouse):
							t.Fatalf("!notLineCompare(exists, got.Warehouse)")
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
