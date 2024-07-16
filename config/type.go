package config

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type AppConfig struct {
	Name            string `mapstructure:"HOST_IP" validate:"required"`
	Env             string `mapstructure:"ENV" validate:"required"`
	Port            int    `mapstructure:"APP_PORT" validate:"required"`
	AccessTokenExp  int    `mapstructure:"ACCESS_TOKEN_EXP" validate:"required"`
	RefreshTokenExp int    `mapstructure:"REFRESH_TOKEN_EXP" validate:"required"`
	JwtSecret       string `mapstructure:"JWT_SECRET" validate:"required"`
	DbType          string `mapstructure:"DB_TYPE" validate:"required"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     string `mapstructure:"DB_PORT" validate:"required"`
	Username string `mapstructure:"DB_USER" validate:"required"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
	SSLMode  string `mapstructure:"SSL_MODE" validate:"required"`
}

type RedisConfig struct {
	Host     string `mapstructure:"CACHE_HOST" validate:"required"`
	Password string `mapstructure:"CACHE_PASSWORD" validate:"required"`
	DB       int    `mapstructure:"CACHE_DB" validate:"required"`
}
