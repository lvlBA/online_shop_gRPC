package user

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/lvlBA/online_shop/pkg/api/v1"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

const (
	firstName = "Daniil"
	lastName  = "Petrushin"
	age       = 34
	login     = "Polter"
	pass      = "Ae112233!"
)

type userResponse interface {
	GetUser() *api.User
}

func compareUserResponse(want, got userResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	return compareUser(want.GetUser(), got.GetUser())
}

func compareUser(want, got *api.User) bool {
	return (want.Id == got.Id) &&
		(want.FirstName == got.FirstName) &&
		(want.LastName == got.LastName) &&
		(want.Login == got.Login) &&
		(want.Age == got.Age) &&
		(want.Sex == got.Sex)
}

func invalidFatal(t *testing.T, want, got interface{}) {
	wantB, _ := json.MarshalIndent(want, "", "   ")
	wantA, _ := json.MarshalIndent(got, "", "   ")
	t.Fatalf("invalid response: \nwant:%s\ngot:%s", wantB, wantA)
}

func compareListUsers(want, got *api.ListUsersResponse) bool {
	if want == nil {
		if got != nil {
			return false
		}
		return true
	}
	if len(want.Users) != len(got.Users) {
		return false
	}
	for i := 0; i < len(want.Users); i++ {
		if !compareUser(want.Users[i], got.Users[i]) {
			return false
		}
	}

	return true
}

func notLineCompare(exists, got []*api.User) bool {
	for i := range got {
		found := false
		for y := range exists {
			if compareUser(exists[y], got[i]) {
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

func Test_User(t *testing.T) {
	ctx := context.Background()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("failed to parse config: %s", err)
	}

	cli, err := New(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to create user: %s", err)
	}

	t.Run("create_user_success", func(t *testing.T) {
		// define

		want := &api.CreateUserResponse{User: &api.User{
			Id:        "",
			FirstName: firstName,
			LastName:  lastName,
			Age:       uint64(age),
			Sex:       api.Sex_SexMale,
			Login:     login,
		}}
		// check

		got, err := cli.CreateUser(ctx, &api.CreateUserRequest{
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
		want.User.Id = got.User.Id
		defer func() {
			if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: got.User.Id}); err != nil {
				t.Fatal(err)
			}
		}()
		if !compareUserResponse(want, got) {
			invalidFatal(t, want, got)
		}
	})
	t.Run("create_user_failed", func(t *testing.T) {
		t.Run("user_exists", func(t *testing.T) {
			// define

			resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
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
				if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: resp.User.Id}); err != nil {
					t.Fatal(err)
				}
			}()
			// check

			resp2, err := cli.CreateUser(ctx, &api.CreateUserRequest{
				FirstName: firstName,
				LastName:  lastName,
				Age:       uint64(age),
				Sex:       api.Sex_SexMale,
				Login:     login,
				Pass:      pass,
			})
			if err == nil {
				defer func() {
					if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: resp2.User.Id}); err != nil {
						t.Fatal(err)
					}
				}()
				t.Fatalf("double user creation")
			}

			assert.Equalf(t, status.Code(err), codes.AlreadyExists,
				"invalid code response: want(%s) got(%s)", codes.AlreadyExists, status.Code(err))
			assert.Nil(t, resp2)
		})
		t.Run("empty", func(t *testing.T) {

			t.Run("first_Name", func(t *testing.T) {
				// check

				resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: "",
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("last_Name", func(t *testing.T) {
				// check

				resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  "",
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("age", func(t *testing.T) {
				// check

				resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       0,
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("login", func(t *testing.T) {
				// check

				resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     "",
					Pass:      pass,
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("pass", func(t *testing.T) {
				// check

				resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      "",
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("sex", func(t *testing.T) {
				// check

				resp, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       uint64(age),
					Sex:       0,
					Login:     login,
					Pass:      pass,
				})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})
	})
	t.Run("get_success", func(t *testing.T) {
		// define

		user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Sex:       api.Sex_SexMale,
			Login:     login,
			Pass:      pass,
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &api.GetUserResponse{User: &api.User{
			Id:        user.User.Id,
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Sex:       api.Sex_SexMale,
			Login:     login,
		}}
		// check

		got, err := cli.GetUser(ctx, &api.GetUserRequest{Id: user.User.Id})

		if err != nil {
			t.Fatal(err)
		}
		if !compareUserResponse(want, got) {
			invalidFatal(t, want, got)
		}
		defer func() {
			if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: got.User.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("get_failed", func(t *testing.T) {
		t.Run("user_doesn't_exist", func(t *testing.T) {
			t.Run("user_is_bad", func(t *testing.T) {
				// check

				resp, err := cli.GetUser(ctx, &api.GetUserRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("user_is_empty", func(t *testing.T) {
				// check

				resp, err := cli.GetUser(ctx, &api.GetUserRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})
	})
	t.Run("delete_success", func(t *testing.T) {
		// define

		user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Sex:       api.Sex_SexMale,
			Login:     login,
			Pass:      pass,
		})
		if err != nil {
			t.Fatal(err)
		}

		// check

		_, err = cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id})

		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)

	})
	t.Run("delete_failed", func(t *testing.T) {
		t.Run("User_doesn't_exist", func(t *testing.T) {
			t.Run("user_is_bad", func(t *testing.T) {
				// check

				resp, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: "123456"})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
			t.Run("userId_is_empty", func(t *testing.T) {
				// check

				resp, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: ""})

				assert.Equalf(t, status.Code(err), codes.InvalidArgument,
					"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
				assert.Nil(t, resp)
			})
		})
	})
	t.Run("list_user_pagination", func(t *testing.T) {
		t.Run("success_pagination", func(t *testing.T) {
			t.Run("without_pagination", func(t *testing.T) {
				// define

				user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
					FirstName: firstName,
					LastName:  lastName,
					Age:       age,
					Sex:       api.Sex_SexMale,
					Login:     login,
					Pass:      pass,
				})
				if err != nil {
					t.Fatal(err)
				}
				want := &api.ListUsersResponse{Users: []*api.User{
					{
						Id:        user.User.Id,
						FirstName: user.User.FirstName,
						LastName:  user.User.LastName,
						Age:       user.User.Age,
						Sex:       user.User.Sex,
						Login:     user.User.Login,
					},
				}}

				// check

				got, err := cli.ListUsers(ctx, &api.ListUsersRequest{Pagination: nil})
				if err != nil {
					t.Fatal(err)
				}
				if !compareListUsers(want, got) {
					wantB, _ := json.MarshalIndent(want, "", "   ")
					gotB, _ := json.MarshalIndent(got, "", "   ")
					t.Fatalf("invalid response:\nwant:%s\ngot:%s", wantB, gotB)
				}
				defer func() {
					if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
						t.Fatal(err)
					}
				}()
			})
			t.Run("with_filter", func(t *testing.T) {
				// define

				exists := make([]*api.User, 0, 20)
				for i := 0; i < len(exists); i++ {
					got, err := cli.CreateUser(ctx, &api.CreateUserRequest{
						FirstName: firstName + fmt.Sprint(i),
						LastName:  lastName + fmt.Sprint(i),
						Age:       uint64(age + i),
						Sex:       api.Sex_SexMale,
						Login:     login + fmt.Sprint(i),
						Pass:      pass + fmt.Sprint(i),
					})
					if err != nil {
						t.Fatal(err)
					}
					exists = append(exists, got.User)
					defer func() {
						if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: got.User.Id}); err != nil {
							t.Fatal(err)
						}
					}()
				}
				// checks

				for limit := 0; limit <= len(exists)+1; limit++ {
					t.Run(fmt.Sprintf("limit %d", limit), func(t *testing.T) {
						got, err := cli.ListUsers(ctx, &api.ListUsersRequest{
							Pagination: &v1.Pagination{
								Page:  1,
								Limit: uint64(limit),
							},
						})
						switch {
						case err != nil:
							t.Fatal(err)
						case limit >= len(exists) && len(got.Users) < len(exists):
							t.Fatalf(" limit >= len(exists) && len(got.Users) < len(exists)")
						case limit < len(exists) && len(got.Users) != limit:
							t.Fatalf(" limit < len(exists) && len(got.Users) != limit")
						case !notLineCompare(exists, got.Users):
							t.Fatalf("!notLineCompare(exists, got.Users)")
						}
					})
				}
			})
		})
	})
	t.Run("change_user_pass_success", func(t *testing.T) {
		// define

		user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Sex:       api.Sex_SexMale,
			Login:     login,
			Pass:      pass,
		})
		if err != nil {
			t.Fatal(err)
		}
		// check

		refreshPass := "Ae11223344!"
		_, err = cli.ChangePass(ctx, &api.ChangePassRequest{
			Id:      user.User.Id,
			OldPass: pass,
			NewPass: refreshPass,
		})
		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)
		defer func() {
			if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
				t.Fatal(err)
			}
		}()
	})
	t.Run("change_pass_failed", func(t *testing.T) {
		t.Run("wrong_user_id", func(t *testing.T) {
			// define

			user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
				FirstName: firstName,
				LastName:  lastName,
				Age:       age,
				Sex:       api.Sex_SexMale,
				Login:     login,
				Pass:      pass,
			})
			if err != nil {
				t.Fatal(err)
			}
			// check

			refreshPass := "Ae11223344!"
			resp, err := cli.ChangePass(ctx, &api.ChangePassRequest{
				Id:      "",
				OldPass: pass,
				NewPass: refreshPass,
			})
			if err == nil {
				t.Fatal(err)
			}
			assert.Equalf(t, status.Code(err), codes.InvalidArgument,
				"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)

			defer func() {
				if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
					t.Fatal(err)
				}
			}()
		})
		t.Run("wrong_pass", func(t *testing.T) {
			// define

			user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
				FirstName: firstName,
				LastName:  lastName,
				Age:       age,
				Sex:       api.Sex_SexMale,
				Login:     login,
				Pass:      pass,
			})
			if err != nil {
				t.Fatal(err)
			}

			// check

			refreshPass := "Ae11223344!"
			resp, err := cli.ChangePass(ctx, &api.ChangePassRequest{
				Id:      user.User.Id,
				OldPass: "Ae112233",
				NewPass: refreshPass,
			})
			if err == nil {
				t.Fatal(err)
			}
			assert.Equalf(t, status.Code(err), codes.Internal,
				"invalid code response: want(%s) got(%s)", codes.Internal, status.Code(err))
			assert.Nil(t, resp)

			defer func() {
				if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
					t.Fatal(err)
				}
			}()
		})
		t.Run("empty_2nd_pass", func(t *testing.T) {
			// define

			user, err := cli.CreateUser(ctx, &api.CreateUserRequest{
				FirstName: firstName,
				LastName:  lastName,
				Age:       age,
				Sex:       api.Sex_SexMale,
				Login:     login,
				Pass:      pass,
			})
			if err != nil {
				t.Fatal(err)
			}
			// check

			_, err = cli.ChangePass(ctx, &api.ChangePassRequest{
				Id:      user.User.Id,
				OldPass: pass,
				NewPass: "",
			})
			if err == nil {
				t.Fatal(err)
			}
			assert.Equalf(t, status.Code(err), codes.InvalidArgument,
				"invalid code response: want(%s) got(%s)", codes.InvalidArgument, status.Code(err))

			defer func() {
				if _, err := cli.DeleteUser(ctx, &api.DeleteUserRequest{Id: user.User.Id}); err != nil {
					t.Fatal(err)
				}
			}()
		})
	})
}
