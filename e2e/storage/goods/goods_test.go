package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	v1 "github.com/lvlBA/online_shop/pkg/api/v1"
	passportapi "github.com/lvlBA/online_shop/pkg/passport/v1"
	passportclient "github.com/lvlBA/online_shop/pkg/passport_client"
	storageapi "github.com/lvlBA/online_shop/pkg/storage/v1"
	storageClient "github.com/lvlBA/online_shop/pkg/storage_client"
)

type GoodsResponse interface {
	GetGoods() *storageapi.Goods
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareGoodsResp(want, got GoodsResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareGoods(want.GetGoods(), got.GetGoods())
}

func compareGoods(want, got *storageapi.Goods) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListGoods(want, got *storageapi.ListGoodsResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Goods) != len(got.Goods) {
		return false
	}
	for i := 0; i < len(want.Goods); i++ {
		if !compareGoods(want.Goods[i], got.Goods[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*storageapi.Goods) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareGoods(exists[y], got[i]) {
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

func Test_goods(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("failed to parse config: %s", err)
	}

	intrUserMeta := grpcinterceptors.NewSetUserMeta()

	pCli, err := passportclient.New(ctx, &passportclient.Config{
		Addr: cfg.PassportAddr,
	}, intrUserMeta.GrpcInterceptor)
	if err != nil {
		t.Fatalf("failed to create passport: %s", err)
	}

	sCli, err := storageClient.New(ctx, &storageClient.Config{
		Addr: cfg.Addr,
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

	resourceCreateGoods, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.GoodsService/CreateGoods",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateGoods': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateGoods.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateGoods': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateGoods.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateGoods.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteGoods, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.GoodsService/DeleteGoods",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteGoods': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteGoods.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteGoods': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteGoods.Resource.Id,
	}); err != nil {
		t.Fatalf("failed to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteGoods.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceGetGoods, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.GoodsService/GetGoods",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetGoods': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceGetGoods.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'GetGoods': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceGetGoods.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceGetGoods.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListGoods, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.GoodsService/ListGoods",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListGoods': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListGoods.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListGoods': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListGoods.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListGoods.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			goodsName := "test_goods1"
			want := &storageapi.CreateGoodsResponse{Goods: &storageapi.Goods{
				Id:     "",
				Name:   goodsName,
				Weight: 2,
				Length: 2,
				Width:  2,
				Height: 2,
				Price:  23.5,
			}}

			// check
			got, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
				Name:   goodsName,
				Weight: 2,
				Length: 2,
				Width:  2,
				Height: 2,
				Price:  23.5,
			})
			if err != nil {
				t.Fatalf("failed to create goods: %s", err)
			}
			want.Goods.Id = got.Goods.Id
			defer func() {
				if _, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: got.Goods.Id}); err != nil {
					t.Fatalf("failed to delete goods: %s", err)
				}
			}()

			if !compareGoodsResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("goods exists", func(t *testing.T) {
				goodsName := "test_goods1"
				resp, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
					Name:   goodsName,
					Weight: 2,
					Length: 2,
					Width:  2,
					Height: 2,
					Price:  200,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: resp.Goods.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
					Name:   goodsName,
					Weight: 2,
					Length: 2,
					Width:  2,
					Height: 2,
					Price:  200,
				})
				if err == nil {
					defer func() {
						if _, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: resp2.Goods.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double goods creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
							Name:   "",
							Weight: 0,
							Length: 0,
							Width:  0,
							Height: 0,
							Price:  0,
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
		goods, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
			Name:   "test",
			Weight: 2,
			Length: 2,
			Width:  2,
			Height: 2,
			Price:  200,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &storageapi.GetGoodsResponse{Goods: &storageapi.Goods{
			Id:     goods.Goods.Id,
			Name:   goods.Goods.Name,
			Weight: goods.Goods.Weight,
			Length: goods.Goods.Length,
			Width:  goods.Goods.Width,
			Height: goods.Goods.Height,
			Price:  goods.Goods.Price,
		}}

		// check
		got, err := sCli.GetGoods(ctx, &storageapi.GetGoodsRequest{
			Id:   goods.Goods.Id,
			Name: goods.Goods.Name,
		})

		if err != nil {
			t.Fatal(err)
		}
		if !compareGoodsResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: got.Goods.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("goods_doesn't exist", func(t *testing.T) {
			t.Run("goods_is_bad", func(t *testing.T) {
				resp, err := sCli.GetGoods(ctx, &storageapi.GetGoodsRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("goods_is_empty", func(t *testing.T) {
				resp, err := sCli.GetGoods(ctx, &storageapi.GetGoodsRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		goods, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
			Name:   "test",
			Weight: 123,
			Length: 256,
			Width:  222,
			Height: 118,
			Price:  0.95,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: goods.Goods.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("goods_doesn't exist", func(t *testing.T) {
			t.Run("goods_is_bad", func(t *testing.T) {
				resp, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("goods_is_empty", func(t *testing.T) {
				resp, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("list_goods_Pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define

				goods, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
					Name:   "test",
					Weight: 1,
					Length: 2,
					Width:  3,
					Height: 4,
					Price:  5,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &storageapi.ListGoodsResponse{Goods: []*storageapi.Goods{
					{
						Id:     goods.Goods.Id,
						Name:   "test",
						Weight: 1,
						Length: 2,
						Width:  3,
						Height: 4,
						Price:  5,
					},
				}}

				// check
				got, err := sCli.ListGoods(ctx, &storageapi.ListGoodsRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListGoods(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: goods.Goods.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*storageapi.Goods, 0, 20)
				for i := 0; i < len(exists); i++ {
					goodsName := fmt.Sprint(i + 252)
					got, err := sCli.CreateGoods(ctx, &storageapi.CreateGoodsRequest{
						Name:   goodsName,
						Weight: 1,
						Length: 2,
						Width:  3,
						Height: 4,
						Price:  5,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Goods)
					defer func() {
						if _, err := sCli.DeleteGoods(ctx, &storageapi.DeleteGoodsRequest{Id: got.Goods.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := sCli.ListGoods(ctx, &storageapi.ListGoodsRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Goods) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Goods) < len(exists)")
						case limit < len(exists) && len(got.Goods) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Goods) != limit")
						case !notLineCompare(exists, got.Goods):
							t.Fatalf("!notLineCompare(exists, got.Goods)")
						}
					})
				}
			})

		})

	})
}
