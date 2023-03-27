package region

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v7"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	v1 "github.com/lvlBA/online_shop/pkg/api/v1"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
	managementclient "github.com/lvlBA/online_shop/pkg/management_client"
	passportapi "github.com/lvlBA/online_shop/pkg/passport/v1"
	passportclient "github.com/lvlBA/online_shop/pkg/passport_client"
)

type regionResponse interface {
	GetRegion() *api.Region
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareRegionResp(want, got regionResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareRegion(want.GetRegion(), got.GetRegion())
}

func compareRegion(want, got *api.Region) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListRegions(want, got *api.ListRegionsResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Region) != len(got.Region) {
		return false
	}
	for i := 0; i < len(want.Region); i++ {
		if !compareRegion(want.Region[i], got.Region[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.Region) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareRegion(exists[y], got[i]) {
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

func Test_region(t *testing.T) {
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

	ctx = context.WithValue(ctx, grpcinterceptors.TokenKey, token.Token)

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

	resourceGetRegion, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.RegionService/GetRegion",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetRegion': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceGetRegion.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'GetRegion': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceGetRegion.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceGetRegion.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListRegion, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.RegionService/ListRegion",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListRegion': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListRegion.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListRegion': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListRegion.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListRegion.Resource.Id,
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

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			regionName := "test_region1"
			want := &api.CreateRegionResponse{Region: &api.Region{
				Id:   "",
				Name: regionName,
			}}

			// check
			got, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
				Name:   regionName,
				SiteId: site.Site.Id,
			})
			if err != nil {
				t.Fatalf("failed to create region: %s", err)
			}
			want.Region.Id = got.Region.Id
			defer func() {
				if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: got.Region.Id}); err != nil {
					t.Fatalf("failed to delete region: %s", err)
				}
			}()

			if !compareRegionResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("region exists", func(t *testing.T) {
				regionName := "test_region1"
				resp, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
					Name:   regionName,
					SiteId: site.Site.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: resp.Region.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
					Name:   regionName,
					SiteId: site.Site.Id,
				})
				if err == nil {
					defer func() {
						if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: resp2.Region.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double region creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
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
		region, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
			Name:   "test",
			SiteId: site.Site.Id,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetRegionResponse{Region: &api.Region{
			Id:   region.Region.Id,
			Name: "test",
		}}

		// check
		got, err := mCli.GetRegion(ctx, &api.GetRegionRequest{Id: region.Region.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareRegionResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: got.Region.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("region_doesn't exist", func(t *testing.T) {
			t.Run("region_is_bad", func(t *testing.T) {
				resp, err := mCli.GetRegion(ctx, &api.GetRegionRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("region_is_empty", func(t *testing.T) {
				resp, err := mCli.GetRegion(ctx, &api.GetRegionRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		region, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
			Name:   "test",
			SiteId: site.Site.Id,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: region.Region.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("region_doesn't exist", func(t *testing.T) {
			t.Run("region_is_bad", func(t *testing.T) {
				resp, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("region_is_empty", func(t *testing.T) {
				resp, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("listRegion_Pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define
				region, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
					Name:   "test",
					SiteId: site.Site.Id,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListRegionsResponse{Region: []*api.Region{
					{
						Id:   region.Region.Id,
						Name: "test",
					},
				}}

				// check
				got, err := mCli.ListRegion(ctx, &api.ListRegionsRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListRegions(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: region.Region.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.Region, 0, 20)
				for i := 0; i < len(exists); i++ {
					regionName := fmt.Sprint(i + 252)
					got, err := mCli.CreateRegion(ctx, &api.CreateRegionRequest{
						Name:   regionName,
						SiteId: site.Site.Id,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Region)
					defer func() {
						if _, err := mCli.DeleteRegion(ctx, &api.DeleteRegionRequest{Id: got.Region.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := mCli.ListRegion(ctx, &api.ListRegionsRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Region) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Regions) < len(exists)")
						case limit < len(exists) && len(got.Region) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Regions) != limit")
						case !notLineCompare(exists, got.Region):
							t.Fatalf("!notLineCompare(exists, got.Regions)")
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
