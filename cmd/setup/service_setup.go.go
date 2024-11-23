package setup

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/application/utils"
)

type Services struct {
	UserService *user.UserService
	AuthService *auth.AuthService
}

func SetupServices(repositories *Repositories, appConfig config.AppConfig) *Services {
	userService := user.NewUserService(repositories.UserRepository)
	authService := auth.NewAuthService(repositories.AuthRepository, userService, repositories.CacheRepository, appConfig, utils.DefaultCheckHash)

	return &Services{
		UserService: &userService,
		AuthService: &authService,
	}
}
