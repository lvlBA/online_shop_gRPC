package passport

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/caarlos0/env/v7"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
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

	t.Run("Create_user_success", func(t *testing.T) {
		// define
		firstName := "Daniil"
		lastName := "Petrushin"
		age := 34
		login := "Polter"
		pass := "Ae112233!"

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
}
