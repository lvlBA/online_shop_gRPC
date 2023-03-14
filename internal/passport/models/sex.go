package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type Sex api.Sex

func (s *Sex) String() string {
	return api.Sex(*s).String()
}

func (s Sex) Value() (driver.Value, error) {
	str := s.String()
	if str == "" {
		return nil, errors.New("invalid value")
	}

	return str, nil
}

func (s *Sex) Scan(src interface{}) error {
	data, ok := src.(string)
	if !ok {
		return errors.New("failed to type assertion to int64")
	}

	value, ok := api.Sex_value[data]
	if !ok {
		return fmt.Errorf("unknown data: (%s)", data)
	}

	*s = Sex(value)

	return nil
}
