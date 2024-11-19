package utils

import (
	"DDD-HEX/config"
	"fmt"
	"net/url"
)

func GeneratePostgresConnectionString(config config.DatabaseConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		url.QueryEscape(config.Username),
		url.QueryEscape(config.Password),
		url.QueryEscape(config.Host),
		url.QueryEscape(config.Port),
		url.QueryEscape(config.Name),
		config.SSLMode)

}
