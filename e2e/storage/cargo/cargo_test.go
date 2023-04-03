package storage

import (
	"context"
	"encoding/json"
	"fmt"
	passportapi "github.com/lvlBA/online_shop/pkg/passport/v1"
	passportclient "github.com/lvlBA/online_shop/pkg/passport_client"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	v1 "github.com/lvlBA/online_shop/pkg/api/v1"
	storageapi "github.com/lvlBA/online_shop/pkg/storage/v1"
	storageClient "github.com/lvlBA/online_shop/pkg/storage_client"
)

type CarrierResponse interface {
	GetCarrier() *storageapi.Carrier
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareCarrierResp(want, got CarrierResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareCarrier(want.GetCarrier(), got.GetCarrier())
}

func compareCarrier(want, got *storageapi.Carrier) bool {
	return want.Id == got.Id && want.Name == got.Name
}

func compareListCarriers(want, got *storageapi.ListCarrierResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Carrier) != len(got.Carrier) {
		return false
	}
	for i := 0; i < len(want.Carrier); i++ {
		if !compareCarrier(want.Carrier[i], got.Carrier[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*storageapi.Carrier) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareCarrier(exists[y], got[i]) {
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

func Test_cargo(t *testing.T) {
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

	resourceCreateCargo, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.CargoService/CreateCarrier",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'CreateCarrier': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceCreateCargo.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'CreateCarrier': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceCreateCargo.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceCreateCargo.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceDeleteCargo, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.CargoService/DeleteCarrier",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'DeleteCarrier': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceDeleteCargo.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'DeleteCarrier': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceDeleteCargo.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceDeleteCargo.Resource.Id,
		}); err != nil {
			t.Fatalf("faield to delete access to resource: %s", err)
		}
	}()

	resourceGetCargo, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.CargoService/GetCarrier",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'GetCarrier': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceGetCargo.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'GetCarrier': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceGetCargo.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceGetCargo.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	resourceListCargo, err := pCli.CreateResource(ctx, &passportapi.CreateResourceRequest{
		Urn: "/online_shop.storage.v1.CargoService/ListCarriers",
	})
	if err != nil {
		t.Fatalf("failed to create resource 'ListCarriers': %s", err)
	}
	defer func() {
		if _, err := pCli.DeleteResource(ctx, &passportapi.DeleteResourceRequest{Id: resourceListCargo.Resource.Id}); err != nil {
			t.Fatalf("failed to delete resource 'ListCarriers': %s", err)
		}
	}()

	if _, err = pCli.SetUserAccess(ctx, &passportapi.SetUserAccessRequest{
		UserId:     user.User.Id,
		ResourceId: resourceListCargo.Resource.Id,
	}); err != nil {
		t.Fatalf("faield to set access to resource: %s", err)
	}
	defer func() {
		if _, err = pCli.DeleteUserAccess(ctx, &passportapi.DeleteUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resourceListCargo.Resource.Id,
		}); err != nil {
			t.Fatalf("failed to delete access to resource: %s", err)
		}
	}()

	t.Run("create", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define
			cargoName := "test_cargo1"
			want := &storageapi.CreateCarrierResponse{Carrier: &storageapi.Carrier{
				Id:           "",
				Name:         cargoName,
				Capacity:     200,
				Price:        200,
				Availability: true,
			}}

			// check
			got, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
				Name:         cargoName,
				Capacity:     200,
				Price:        200,
				Availability: true,
			})
			if err != nil {
				t.Fatalf("failed to create carrier: %s", err)
			}
			want.Carrier.Id = got.Carrier.Id
			defer func() {
				if _, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: got.Carrier.Id}); err != nil {
					t.Fatalf("failed to delete carrier: %s", err)
				}
			}()

			if !compareCarrierResp(want, got) {
				invalidFatal(t, want, got)
			}
		})

		t.Run("failed", func(t *testing.T) {

			t.Run("carrier exists", func(t *testing.T) {
				cargoName := "test_carrier1"
				resp, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
					Name:         cargoName,
					Capacity:     200,
					Price:        200,
					Availability: true,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: resp.Carrier.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				resp2, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
					Name:         cargoName,
					Capacity:     200,
					Price:        200,
					Availability: true,
				})
				if err == nil {
					defer func() {
						if _, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: resp2.Carrier.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					t.Fatalf("double carrier creation")
				}

				assert.Equalf(t, status.Code(err), codes.AlreadyExists,
					"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
				assert.Nil(t, resp2)
			})

			t.Run("invalid argument", func(t *testing.T) {

				t.Run("empty", func(t *testing.T) {

					t.Run("name", func(t *testing.T) {
						resp, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
							Name:         "",
							Capacity:     0,
							Price:        0,
							Availability: false,
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
		carrier, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
			Name:         "test",
			Capacity:     200,
			Price:        200,
			Availability: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &storageapi.GetCarrierResponse{Carrier: &storageapi.Carrier{
			Id:           carrier.Carrier.Id,
			Name:         carrier.Carrier.Name,
			Capacity:     carrier.Carrier.Capacity,
			Price:        carrier.Carrier.Price,
			Availability: carrier.Carrier.Availability,
		}}

		// check
		got, err := sCli.GetCarrier(ctx, &storageapi.GetCarrierRequest{
			Id:   carrier.Carrier.Id,
			Name: carrier.Carrier.Name,
		})

		if err != nil {
			t.Fatal(err)
		}
		if !compareCarrierResp(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: got.Carrier.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_Failed", func(t *testing.T) {
		t.Run("carrier_doesn't exist", func(t *testing.T) {
			t.Run("carrier_is_bad", func(t *testing.T) {
				resp, err := sCli.GetCarrier(ctx, &storageapi.GetCarrierRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("carrier_is_empty", func(t *testing.T) {
				resp, err := sCli.GetCarrier(ctx, &storageapi.GetCarrierRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("delete_Success", func(t *testing.T) {
		// define

		carrier, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
			Name:         "test",
			Capacity:     200,
			Price:        200,
			Availability: true,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check
		_, err = sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: carrier.Carrier.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_Failed", func(t *testing.T) {
		t.Run("carrier_doesn't exist", func(t *testing.T) {
			t.Run("carrier_is_bad", func(t *testing.T) {
				resp, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("carrier_is_empty", func(t *testing.T) {
				resp, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})

	})
	t.Run("listCarrier_Pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without pagination", func(t *testing.T) {
				// define

				carrier, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
					Name:         "carrierName",
					Capacity:     200,
					Price:        200,
					Availability: true,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &storageapi.ListCarrierResponse{Carrier: []*storageapi.Carrier{
					{
						Id:           carrier.Carrier.Id,
						Name:         "carrierName",
						Capacity:     200,
						Price:        200,
						Availability: true,
					},
				}}

				// check
				got, err := sCli.ListCarriers(ctx, &storageapi.ListCarrierRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListCarriers(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: carrier.Carrier.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with filter", func(t *testing.T) {
				// define
				exists := make([]*storageapi.Carrier, 0, 20)
				for i := 0; i < len(exists); i++ {
					carrierName := fmt.Sprint(i + 252)
					got, err := sCli.CreateCarrier(ctx, &storageapi.CreateCarrierRequest{
						Name:         carrierName,
						Capacity:     200,
						Price:        200,
						Availability: false,
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.Carrier)
					defer func() {
						if _, err := sCli.DeleteCarrier(ctx, &storageapi.DeleteCarrierRequest{Id: got.Carrier.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}

				// checks
				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := sCli.ListCarriers(ctx, &storageapi.ListCarrierRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Carrier) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Carriers) < len(exists)")
						case limit < len(exists) && len(got.Carrier) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Carriers) != limit")
						case !notLineCompare(exists, got.Carrier):
							t.Fatalf("!notLineCompare(exists, got.Carriers)")
						}
					})
				}
			})

		})

	})
}
