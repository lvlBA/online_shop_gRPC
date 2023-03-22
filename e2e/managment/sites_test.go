package managment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v7"
	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
	managementclient "github.com/lvlBA/online_shop/pkg/management_client"
	passportapi "github.com/lvlBA/online_shop/pkg/passport/v1"
	passportclient "github.com/lvlBA/online_shop/pkg/passport_client"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type siteResponse interface {
	GetSite() *api.Site
}

func compareSiteResp(want, got siteResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareSite(want.GetSite(), got.GetSite())
}

func compareSite(want, got *api.Site) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListSites(want, got *api.ListSitesResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Sites) != len(got.Sites) {
		return false
	}
	for i := 0; i < len(want.Sites); i++ {
		if !compareSite(want.Sites[i], got.Sites[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.Site) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareSite(exists[y], got[i]) {
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

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func Test_sites(t *testing.T) {
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
		t.Fatalf("failed to create site: %s", err)
	}

	pCli, err := passportclient.New(ctx, &passportclient.Config{
		Addr: cfg.PassportAddr,
	}, intrUserMeta.GrpcInterceptor)
	if err != nil {
		t.Fatalf("failed to create site: %s", err)
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

	resourceCheckSite, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.SiteService/GetSite",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetSite': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCheckSite.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteSite': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCheckSite.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCheckSite.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListSite, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.management.v1.SiteService/ListSites",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListSites': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListSite.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListSites': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListSite.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListSite.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			siteName := "test_site1"
			want := &api.CreateSideResponse{Site: &api.Site{
				Id:   "",
				Name: siteName,
			}}

			// check
			got, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
				Name: siteName,
			})
			if err != nil {
				t.Fatalf("failed to create site: %s", err)
			}
			want.Site.Id = got.Site.Id
			defer func() {
				if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: got.Site.Id}); err != nil {
					t.Fatalf("failed to delete site: %s", err)
				}
			}()

			if !compareSiteResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("site exists", func(t *testing.T) {
				siteName := "test_site1"
				resp, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
					Name: siteName,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: resp.Site.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
					Name: siteName,
				})
				if err == nil {
					defer func() {
						if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: resp2.Site.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double site creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
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
		site, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
			Name: "test",
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetSiteResponse{Site: &api.Site{
			Id:   site.Site.Id,
			Name: "test",
		}}

		// check
		got, err := mCli.GetSite(ctx, &api.GetSiteRequest{Id: site.Site.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareSiteResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: got.Site.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("site_doesn't exist", func(t *testing.T) {
			t.Run("site_is_bad", func(t *testing.T) {
				resp, err := mCli.GetSite(ctx, &api.GetSiteRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("site_is_empty", func(t *testing.T) {
				resp, err := mCli.GetSite(ctx, &api.GetSiteRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		site, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
			Name: "test",
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: site.Site.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("site_doesn't exist", func(t *testing.T) {
			t.Run("Site_is_bad", func(t *testing.T) {
				resp, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("Site_is_empty", func(t *testing.T) {
				resp, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("listSite_Pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define
				site, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
					Name: "test",
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListSitesResponse{Sites: []*api.Site{
					{
						Id:   site.Site.Id,
						Name: "test",
					},
				}}

				// check
				got, err := mCli.ListSites(ctx, &api.ListSitesRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListSites(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: site.Site.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.Site, 0, 20)
				for i := 0; i < len(exists); i++ {
					siteName := fmt.Sprint(i + 252)
					got, err := mCli.CreateSite(ctx, &api.CreateSideRequest{
						Name: siteName,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Site)
					defer func() {
						if _, err := mCli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: got.Site.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := mCli.ListSites(ctx, &api.ListSitesRequest{
							Pagination: &api.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Sites) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Sites) < len(exists)")
						case limit < len(exists) && len(got.Sites) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Sites) != limit")
						case !notLineCompare(exists, got.Sites):
							t.Fatalf("!notLineCompare(exists, got.Sites)")
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
