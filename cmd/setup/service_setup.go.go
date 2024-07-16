package setup

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/ports/repository"
)

func SetupServices(userRepo repository.UserRepository, authRepo repository.AuthRepository, cacheRepo repository.CacheRepository, appConfig config.AppConfig) (user.UserService, auth.AuthService) {
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(authRepo, userService, cacheRepo, appConfig)

	return userService, authService
}
