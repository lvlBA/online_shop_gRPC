package managment

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/lvlBA/online_shop/pkg/management/v1"
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

	cli, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to create site: %s", err)
	}

	t.Run("create_success", func(t *testing.T) {
		// define
		siteName := "test_site1"
		want := &api.CreateSideResponse{Site: &api.Site{
			Id:   "",
			Name: siteName,
		}}
		// check
		got, err := cli.CreateSite(ctx, &api.CreateSideRequest{
			Name: siteName,
		})
		if err != nil {
			t.Fatal(err)
		}
		want.Site.Id = got.Site.Id
		defer func() {
			if _, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: got.Site.Id}); err != nil {
				t.Fatal(err)
			}
		}()

		if !compareSiteResp(want, got) {
			invalidFatal(t, want, got)
		}
	})
	t.Run("create_failed", func(t *testing.T) {
		t.Run("site exists", func(t *testing.T) {
			siteName := "test_site1"
			resp, err := cli.CreateSite(ctx, &api.CreateSideRequest{
				Name: siteName,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if _, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: resp.Site.Id}); err != nil {
					t.Fatal(err)
				}
			}()

			resp2, err := cli.CreateSite(ctx, &api.CreateSideRequest{
				Name: siteName,
			})
			if err == nil {
				defer func() {
					if _, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: resp2.Site.Id}); err != nil {
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
					resp, err := cli.CreateSite(ctx, &api.CreateSideRequest{
						Name: "",
					})

					assert.Equalf(t, status.Code(err), codes.InvalidArgument,
						"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
					assert.Nil(t, resp)
				})
			})
		})
	})
	t.Run("get_Success", func(t *testing.T) {
		// define
		site, err := cli.CreateSite(ctx, &api.CreateSideRequest{
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
		got, err := cli.GetSite(ctx, &api.GetSiteRequest{Id: site.Site.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareSiteResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: got.Site.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("site_doesn't exist", func(t *testing.T) {
			t.Run("site_is_bad", func(t *testing.T) {
				resp, err := cli.GetSite(ctx, &api.GetSiteRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("site_is_empty", func(t *testing.T) {
				resp, err := cli.GetSite(ctx, &api.GetSiteRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		site, err := cli.CreateSite(ctx, &api.CreateSideRequest{
			Name: "test",
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: site.Site.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("site_doesn't exist", func(t *testing.T) {
			t.Run("Site_is_bad", func(t *testing.T) {
				resp, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("Site_is_empty", func(t *testing.T) {
				resp, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: ""})

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
				site, err := cli.CreateSite(ctx, &api.CreateSideRequest{
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
				got, err := cli.ListSites(ctx, &api.ListSitesRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListSites(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: site.Site.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.Site, 0, 20)
				for i := 0; i < len(exists); i++ {
					siteName := fmt.Sprint(i + 252)
					got, err := cli.CreateSite(ctx, &api.CreateSideRequest{
						Name: siteName,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Site)
					defer func() {
						if _, err := cli.DeleteSite(ctx, &api.DeleteSiteRequest{Id: got.Site.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := cli.ListSites(ctx, &api.ListSitesRequest{
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
