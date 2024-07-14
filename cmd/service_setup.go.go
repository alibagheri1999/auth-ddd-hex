package cmd

import (
	"DDD-HEX/config"
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/services/product"
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/internal/ports/repository"
)

func setupServices(userRepo repository.UserRepository, authRepo repository.AuthRepository, productRepo repository.ProductRepository, appConfig config.AppConfig) (user.UserService, auth.AuthService, product.ProductService) {
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(authRepo, userRepo, appConfig)
	productService := product.NewProductService(productRepo)

	return userService, authService, productService
}
