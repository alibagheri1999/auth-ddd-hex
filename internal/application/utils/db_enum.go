package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type DbEnum int

const (
	Postgres DbEnum = iota
)

var dbName = []string{"postgres"}

func (r DbEnum) String() string {
	if r < Postgres || r > Postgres {
		return "unknown"
	}
	return dbName[r]
}

func ParseDbEnum(role string) (DbEnum, error) {
	for i, v := range dbName {
		if v == role {
			return DbEnum(i), nil
		}
	}
	return -1, fmt.Errorf("invalid DbEnum value: %s", role)
}

func (r *DbEnum) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type assertion")
	}

	parsedRole, err := ParseDbEnum(str)
	if err != nil {
		return err
	}

	*r = parsedRole
	return nil
}

func (r DbEnum) Value() (driver.Value, error) {
	if r < Postgres || r > Postgres {
		return nil, fmt.Errorf("invalid DbEnum value: %d", r)
	}
	return r.String(), nil
}
