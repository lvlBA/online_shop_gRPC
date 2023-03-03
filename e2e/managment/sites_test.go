package managment

import (
	"context"
	"encoding/json"
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

	t.Run("Success", func(t *testing.T) {
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
	t.Run("failed", func(t *testing.T) {
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

}
