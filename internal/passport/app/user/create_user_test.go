package user

import (
	"testing"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func Test_validateCreateUserReq(t *testing.T) {
	type args struct {
		req *api.CreateUserRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			wantErr: false,
			args: args{req: &api.CreateUserRequest{
				FirstName: "daniel",
				LastName:  "petrushin",
				Age:       34,
				Sex:       api.Sex_SexMale,
				Login:     "Poltergeist",
				Pass:      "Ae112233!",
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreateUserReq(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateCreateUserReq() error = %v, wantErr %v\n", err, tt.wantErr)
			}
		})
	}
}
