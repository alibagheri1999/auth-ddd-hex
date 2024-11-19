package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type CacheEnum int

const (
	Redis CacheEnum = iota
)

var cacheName = []string{"redis"}

func (r CacheEnum) String() string {
	if r < Redis || r > Redis {
		return "unknown"
	}
	return cacheName[r]
}

func ParseCacheEnum(role string) (CacheEnum, error) {
	for i, v := range cacheName {
		if v == role {
			return CacheEnum(i), nil
		}
	}
	return -1, fmt.Errorf("invalid CacheEnum value: %s", role)
}

func (r *CacheEnum) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type assertion")
	}

	parsedRole, err := ParseCacheEnum(str)
	if err != nil {
		return err
	}

	*r = parsedRole
	return nil
}

func (r CacheEnum) Value() (driver.Value, error) {
	if r < Redis || r > Redis {
		return nil, fmt.Errorf("invalid CacheEnum value: %d", r)
	}
	return r.String(), nil
}
