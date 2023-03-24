package location

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

type locationResponse interface {
	GetLocation() *api.Location
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareLocationResp(want, got locationResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareLocation(want.GetLocation(), got.GetLocation())
}

func compareLocation(want, got *api.Location) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListLocations(want, got *api.ListLocationsResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Location) != len(got.Location) {
		return false
	}
	for i := 0; i < len(want.Location); i++ {
		if !compareLocation(want.Location[i], got.Location[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.Location) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareLocation(exists[y], got[i]) {
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

func Test_Location(t *testing.T) {
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

	resourceGetLocation, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.LocationService/GetLocation",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetLocation': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceGetLocation.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'GetLocation': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceGetLocation.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceGetLocation.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListLocation, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.LocationService/ListLocation",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListLocation': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListLocation.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListLocation': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListLocation.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListLocation.Resource.Id,
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

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			locationName := "test_location1"
			want := &api.CreateLocationResponse{Location: &api.Location{
				Id:   "",
				Name: locationName,
			}}

			// check
			got, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
				Name:     locationName,
				SiteId:   site.Site.Id,
				RegionId: region.Region.Id,
			})
			if err != nil {
				t.Fatalf("failed to create location: %s", err)
			}
			want.Location.Id = got.Location.Id
			defer func() {
				if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: got.Location.Id}); err != nil {
					t.Fatalf("failed to delete location: %s", err)
				}
			}()

			if !compareLocationResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("location exists", func(t *testing.T) {
				locationName := "test_location"
				resp, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
					Name:     locationName,
					SiteId:   site.Site.Id,
					RegionId: region.Region.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: resp.Location.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
					Name:     locationName,
					SiteId:   site.Site.Id,
					RegionId: region.Region.Id,
				})
				if err == nil {
					defer func() {
						if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: resp2.Location.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double location creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
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
		location, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
			Name:     "test",
			SiteId:   site.Site.Id,
			RegionId: region.Region.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetLocationResponse{Location: &api.Location{
			Id:   location.Location.Id,
			Name: "test",
		}}

		// check
		got, err := mCli.GetLocation(ctx, &api.GetLocationRequest{Id: location.Location.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareLocationResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: got.Location.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("location_doesn't exist", func(t *testing.T) {
			t.Run("location_is_bad", func(t *testing.T) {
				resp, err := mCli.GetLocation(ctx, &api.GetLocationRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("location_is_empty", func(t *testing.T) {
				resp, err := mCli.GetLocation(ctx, &api.GetLocationRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		location, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
			Name:     "test",
			SiteId:   site.Site.Id,
			RegionId: region.Region.Id,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: location.Location.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("location_doesn't exist", func(t *testing.T) {
			t.Run("location_is_bad", func(t *testing.T) {
				resp, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("location_is_empty", func(t *testing.T) {
				resp, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("listLocation_Pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define
				location, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
					Name:     "test",
					SiteId:   site.Site.Id,
					RegionId: region.Region.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListLocationsResponse{Location: []*api.Location{
					{
						Id:   location.Location.Id,
						Name: "test",
					},
				}}

				// check
				got, err := mCli.ListLocation(ctx, &api.ListLocationsRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListLocations(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: location.Location.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.Location, 0, 20)
				for i := 0; i < len(exists); i++ {
					locationName := fmt.Sprint(i + 252)
					got, err := mCli.CreateLocation(ctx, &api.CreateLocationRequest{
						Name:     locationName,
						SiteId:   site.Site.Id,
						RegionId: region.Region.Id,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Location)
					defer func() {
						if _, err := mCli.DeleteLocation(ctx, &api.DeleteLocationRequest{Id: got.Location.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := mCli.ListLocation(ctx, &api.ListLocationsRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Location) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Location) < len(exists)")
						case limit < len(exists) && len(got.Location) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Location) != limit")
						case !notLineCompare(exists, got.Location):
							t.Fatalf("!notLineCompare(exists, got.Location)")
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
