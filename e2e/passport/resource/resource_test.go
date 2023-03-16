package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

const (
	urn = "Ae112233sdkforNothing"
)

type resourceResponse interface {
	GetResource() *api.Resource
}

func compareResourceResponse(want, got resourceResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareResource(want.GetResource(), got.GetResource())
}

func compareResource(want, got *api.Resource) bool {
	return want.Id == got.Id && want.Urn == got.Urn
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareListResources(want, got *api.ListResourceResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Resource) != len(got.Resource) {
		return false
	}
	for i := 0; i < len(want.Resource); i++ {
		if !compareResource(want.Resource[i], got.Resource[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.Resource) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareResource(exists[y], got[i]) {
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

func Test_resource(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("failed to parse config: %s", err)
	}

	cli, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to create resource: %s", err)
	}
	t.Run("create_resource_success", func(t *testing.T) {
		// define

		want := &api.CreateResourceResponse{Resource: &api.Resource{
			Id:  "",
			Urn: urn,
		}}

		// check

		got, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})

		if err != nil {
			t.Fatal(err)
		}
		want.Resource.Id = got.Resource.Id
		defer func() {
			if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: got.Resource.Id}); err != nil {
				t.Fatal(err)
			}
		}()
		if !compareResourceResponse(want, got) {
			invalidFatal(t, want, got)
		}
	})
	t.Run("create_resource_failed", func(t *testing.T) {
		t.Run("resource exists", func(t *testing.T) {
			// define

			resp, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
				Urn: urn,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resp.Resource.Id}); err != nil {
					t.Fatal(err)
				}
			}()
			// check

			resp2, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
				Urn: urn,
			})
			if err == nil {
				defer func() {
					if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resp2.Resource.Id}); err != nil {
						t.Fatal(err)
					}
				}()
				t.Fatalf("double resource creation")
			}

			assert.Equalf(t, status.Code(err), codes.AlreadyExists,
				"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
			assert.Nil(t, resp2)
		})
		t.Run("empty", func(t *testing.T) {
			t.Run("urn", func(t *testing.T) {
				resp, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
					Urn: "",
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})

		})
	})
	t.Run("get_success", func(t *testing.T) {
		// define
		resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
			Urn: urn,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetResourceResponse{Resource: &api.Resource{
			Id:  resource.Resource.Id,
			Urn: resource.Resource.Urn,
		}}

		// check
		got, err := cli.GetResource(ctx, &api.GetResourceRequest{Id: resource.Resource.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareResourceResponse(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: got.Resource.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_failed", func(t *testing.T) {
		t.Run("user_doesn't exist", func(t *testing.T) {
			t.Run("user_is_bad", func(t *testing.T) {
				resp, err := cli.GetResource(ctx, &api.GetResourceRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("user_is_empty", func(t *testing.T) {
				resp, err := cli.GetResource(ctx, &api.GetResourceRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})
	})
	t.Run("delete_success", func(t *testing.T) {
		// define

		resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
			Urn: urn,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_failed", func(t *testing.T) {
		t.Run("user_doesn't_exist", func(t *testing.T) {
			t.Run("user_is_bad", func(t *testing.T) {
				// check
				resp, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("user_id_is_empty", func(t *testing.T) {
				// check
				resp, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})
	})
	t.Run("list_user_pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define
				resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
					Urn: urn,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListResourceResponse{Resource: []*api.Resource{
					{
						Id:  resource.Resource.Id,
						Urn: resource.Resource.Urn,
					},
				}}

				// check
				got, err := cli.ListResource(ctx, &api.ListResourceRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListResources(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*api.Resource, 0, 20)
				for i := 0; i < len(exists); i++ {
					got, err := cli.CreateResource(ctx, &api.CreateResourceRequest{
						Urn: urn + fmt.Sprint(i),
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Resource)
					defer func() {
						if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: got.Resource.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := cli.ListResource(ctx, &api.ListResourceRequest{
							Pagination: &api.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Resource) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.resources) < len(exists)")
						case limit < len(exists) && len(got.Resource) != limit:
							t.Fatalf(" limit < len(exists) && len(got.resources) != limit")
						case !notLineCompare(exists, got.Resource):
							t.Fatalf("!notLineCompare(exists, got.resources)")
						}
					})
				}
			})
		})
	})
}
