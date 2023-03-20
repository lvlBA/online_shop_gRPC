package auth

import (
	"context"
	grpcinterceptors "github.com/lvlBA/online_shop/internal/grpc_interceptors"
	passportclient "github.com/lvlBA/online_shop/pkg/passport_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/assert"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

const (
	firstName = "Daniil"
	lastName  = "Petrushin"
	age       = 34
	login     = "Polter"
	pass      = "Ae112233!"
	urn       = "Ae112233sdkforNothing"
)

func Test_auth(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("failed to parse config: %s", err)
	}

	intrUserMeta := grpcinterceptors.NewSetUserMeta()

	cli, err := passportclient.New(ctx, &passportclient.Config{
		Addr: cfg.Addr,
	}, intrUserMeta.GrpcInterceptor)
	if err != nil {
		t.Fatalf("failed to create auth: %s", err)
	}

	t.Run("set_user_access_success", func(t *testing.T) {
		// define

		user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
			FirstName: firstName,
			LastName:  lastName,
			Age:       uint64(age),
			Sex:       api.Sex_SexMale,
			Login:     login,
			Pass:      pass,
		})
		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
				t.Fatal(err)
			}
		}()

		resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
				t.Fatal(err)
			}
		}()
		//check

		_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resource.Resource.Id,
		})
		assert.Nil(t, err)
		defer func() {
			if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
				UserId:     user.User.Id,
				ResourceId: resource.Resource.Id,
			}); err != nil {
				t.Fatal(err)
			}
		}()
	})

	t.Run("set_user_access_failed", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			t.Run("userId", func(t *testing.T) {
				// define

				resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
						t.Fatal(err)
					}
				}()
				//check

				_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
					UserId:     "",
					ResourceId: resource.Resource.Id,
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				defer func() {
					if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
						UserId:     "",
						ResourceId: resource.Resource.Id,
					}); err == nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("resourceId", func(t *testing.T) {
				// define

				user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				// check
				_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
					UserId:     user.User.Id,
					ResourceId: "",
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				defer func() {
					if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
						UserId:     user.User.Id,
						ResourceId: "",
					}); err == nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("all_fields_empty", func(t *testing.T) {
				// define
				_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
					UserId:     "",
					ResourceId: "",
				})
				// check

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				defer func() {
					if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
						UserId:     "",
						ResourceId: "",
					}); err == nil {
						t.Fatal(err)
					}
				}()
			})
		})
	})

	t.Run("delete_user_access_success", func(t *testing.T) {
		// define
		user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
			FirstName: firstName,
			LastName:  lastName,
			Age:       uint64(age),
			Sex:       api.Sex_SexMale,
			Login:     login,
			Pass:      pass,
		})
		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
				t.Fatal(err)
			}
		}()

		resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
				t.Fatal(err)
			}
		}()
		// check

		_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
			UserId:     user.User.Id,
			ResourceId: resource.Resource.Id,
		})
		defer func() {
			if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
				UserId:     user.User.Id,
				ResourceId: resource.Resource.Id,
			}); err != nil {
				t.Fatal(err)
			}
		}()
		assert.Nil(t, err)

	})

	t.Run("delete_user_access_failed", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			t.Run("userId", func(t *testing.T) {
				// define
				resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				// check

				_, err = cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
					UserId:     "",
					ResourceId: resource.Resource.Id,
				})
				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))

			})
			t.Run("resourceId", func(t *testing.T) {
				// define

				user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})
				if err != nil {
					t.Fatal(err)
				}

				defer func() {
					if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
						t.Fatal(err)
					}
				}()
				// check

				_, err = cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
					UserId:     user.User.Id,
					ResourceId: "",
				})
				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))

			})
			t.Run("all_fields_empty", func(t *testing.T) {
				// check

				_, err = cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
					UserId:     "",
					ResourceId: "",
				})
				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
			})
		})
	})

	t.Run("get_user_token_success", func(t *testing.T) {
		// define

		user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
			FirstName: firstName,
			LastName:  lastName,
			Age:       uint64(age),
			Sex:       api.Sex_SexMale,
			Login:     login,
			Pass:      pass,
		})
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
				t.Fatal(err)
			}
		}()

		//check
		_, err = cli.GetUserToken(ctx, &api.GetUserTokenRequest{
			Login:    user.User.Login,
			Password: pass,
		})
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if _, err := cli.DeleteUserToken(ctx, &api.DeleteUserTokenRequest{
				Login:    user.User.Login,
				Password: pass,
			}); err != nil {
				t.Fatal(err)
			}
		}()
	})

	t.Run("get_user_token_failed", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			t.Run("userId", func(t *testing.T) {
				// define
				user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				// check
				_, err = cli.GetUserToken(ctx, &api.GetUserTokenRequest{
					Login:    "",
					Password: pass,
				})
				if err == nil {
					t.Fatal(err)
				}
				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
			})
			t.Run("resourceId", func(t *testing.T) {
				// define

				user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})
				if err != nil {
					t.Fatal(err)
				}
				defer func() {
					if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
						t.Fatal(err)
					}
				}()

				//check
				_, err = cli.GetUserToken(ctx, &api.GetUserTokenRequest{
					Login:    user.User.Login,
					Password: "",
				})
				if err == nil {
					t.Fatal(err)
				}
				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
			})
			t.Run("all_fields_empty", func(t *testing.T) {
				//check
				_, err = cli.GetUserToken(ctx, &api.GetUserTokenRequest{
					Login:    "",
					Password: "",
				})
				if err == nil {
					t.Fatal(err)
				}
				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
			})
		})
	})

	t.Run("check_user_access", func(t *testing.T) {

		t.Run("success", func(t *testing.T) {
			// define

			user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
				FirstName: firstName,
				LastName:  lastName,
				Age:       uint64(age),
				Sex:       api.Sex_SexMale,
				Login:     login,
				Pass:      pass,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
					t.Fatal(err)
				}
			}()
			ctx = context.WithValue(ctx, "x-request-id", user.User.Login)

			token, err := cli.GetUserToken(ctx, &api.GetUserTokenRequest{
				Login:    user.User.Login,
				Password: pass,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if _, err := cli.DeleteUserToken(ctx, &api.DeleteUserTokenRequest{
					Login:    user.User.Login,
					Password: pass,
				}); err != nil {
					t.Fatal(err)
				}
			}()
			ctx = context.WithValue(ctx, "token", token.Token)

			resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
					t.Fatal(err)
				}
			}()

			_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
				UserId:     user.User.Id,
				ResourceId: resource.Resource.Id,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
					UserId:     user.User.Id,
					ResourceId: resource.Resource.Id,
				}); err != nil {
					t.Fatal(err)
				}
			}()

			// check
			_, err = cli.CheckUserAccess(ctx, &api.CheckUserAccessRequest{
				ResourceId: resource.Resource.Id,
			})
			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("failed", func(t *testing.T) {
			t.Run("token empty", func(t *testing.T) {
				t.Run("token", func(t *testing.T) {
					// define
					user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
						FirstName: firstName,
						LastName:  lastName,
						Age:       uint64(age),
						Sex:       api.Sex_SexMale,
						Login:     login,
						Pass:      pass,
					})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					ctx = context.WithValue(ctx, "x-request-id", user.User.Login)

					resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
							t.Fatal(err)
						}
					}()

					if _, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
						UserId:     user.User.Id,
						ResourceId: resource.Resource.Id,
					}); err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
							UserId:     user.User.Id,
							ResourceId: resource.Resource.Id,
						}); err != nil {
							t.Fatal(err)
						}
					}()

					//check
					if _, err = cli.CheckUserAccess(ctx, &api.CheckUserAccessRequest{
						ResourceId: resource.Resource.Id,
					}); err == nil {
						t.Fatal(err)
					}
					assert.Equalf(t, status.Code(err), codes.Unauthenticated, "invalid code response: want(%s) got(%s)", codes.Unauthenticated, status.Code(err))
				})

				t.Run("resourceId", func(t *testing.T) {
					// define

					user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
						FirstName: firstName,
						LastName:  lastName,
						Age:       uint64(age),
						Sex:       api.Sex_SexMale,
						Login:     login,
						Pass:      pass,
					})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
							t.Fatal(err)
						}
					}()
					ctx = context.WithValue(ctx, "x-request-id", user.User.Login)

					token, err := cli.GetUserToken(ctx, &api.GetUserTokenRequest{
						Login:    user.User.Login,
						Password: pass,
					})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteUserToken(ctx, &api.DeleteUserTokenRequest{
							Login:    user.User.Login,
							Password: pass,
						}); err != nil {
							t.Fatal(err)
						}
					}()
					ctx = context.WithValue(ctx, "token", token.Token)

					resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
							t.Fatal(err)
						}
					}()

					_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
						UserId:     user.User.Id,
						ResourceId: resource.Resource.Id,
					})
					assert.Nil(t, err)
					defer func() {
						if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
							UserId:     user.User.Id,
							ResourceId: resource.Resource.Id,
						}); err != nil {
							t.Fatal(err)
						}
					}()
					//check

					_, err = cli.CheckUserAccess(ctx, &api.CheckUserAccessRequest{
						ResourceId: "",
					})
					if err == nil {
						t.Fatal(err)
					}
					assert.Equalf(t, status.Code(err), codes.InvalidArgument,
						"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				})

				t.Run("all_fields_empty", func(t *testing.T) {
					// define

					user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
						FirstName: firstName,
						LastName:  lastName,
						Age:       uint64(age),
						Sex:       api.Sex_SexMale,
						Login:     login,
						Pass:      pass,
					})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
							t.Fatal(err)
						}
					}()

					_, err = cli.GetUserToken(ctx, &api.GetUserTokenRequest{
						Login:    user.User.Login,
						Password: pass,
					})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteUserToken(ctx, &api.DeleteUserTokenRequest{
							Login:    user.User.Login,
							Password: pass,
						}); err != nil {
							t.Fatal(err)
						}
					}()

					resource, err := cli.CreateResource(ctx, &api.CreateResourceRequest{Urn: urn})
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						if _, err := cli.DeleteResource(ctx, &api.DeleteResourceRequest{Id: resource.Resource.Id}); err != nil {
							t.Fatal(err)
						}
					}()

					_, err = cli.SetUserAccess(ctx, &api.SetUserAccessRequest{
						UserId:     user.User.Id,
						ResourceId: resource.Resource.Id,
					})
					assert.Nil(t, err)
					defer func() {
						if _, err := cli.DeleteUserAccess(ctx, &api.DeleteUserAccessRequest{
							UserId:     user.User.Id,
							ResourceId: resource.Resource.Id,
						}); err != nil {
							t.Fatal(err)
						}
					}()

					//check

					_, err = cli.CheckUserAccess(ctx, &api.CheckUserAccessRequest{
						ResourceId: "",
					})
					if err == nil {
						t.Fatal(err)
					}
					assert.Equalf(t, status.Code(err), codes.InvalidArgument,
						"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				})
			})
		})
	})
}
