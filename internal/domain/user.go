package domain

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type UserEntity struct {
	ID             int            `db:"id"`
	Name           string         `db:"name"`
	Email          string         `db:"email"`
	Password       sql.NullString `db:"password"`
	Rule           UserRole       `db:"rule"`
	PhoneNumber    sql.NullString `db:"phone_number"`
	TowFAActivated bool           `db:"two_fa_activated"`
	Status         string         `db:"status"`
	DateCreated    time.Time      `db:"date_created"`
}

type UserRole int

const (
	Admin UserRole = iota
	User
)

var userRoles = []string{"admin", "user"}

func (r UserRole) String() string {
	if r < Admin || r > User {
		return "unknown"
	}
	return userRoles[r]
}

func ParseUserRole(role string) (UserRole, error) {
	for i, v := range userRoles {
		if v == role {
			return UserRole(i), nil
		}
	}
	return -1, fmt.Errorf("invalid UserRole value: %s", role)
}

func (r *UserRole) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type assertion")
	}

	parsedRole, err := ParseUserRole(str)
	if err != nil {
		return err
	}

	*r = parsedRole
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	if r < Admin || r > User {
		return nil, fmt.Errorf("invalid UserRole value: %d", r)
	}
	return r.String(), nil
}
