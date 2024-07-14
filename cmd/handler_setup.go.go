package cmd

import (
	Handler "DDD-HEX/internal/adapters/http"
	"DDD-HEX/internal/application/services/auth"
	"DDD-HEX/internal/application/services/product"
	"DDD-HEX/internal/application/services/user"
	"DDD-HEX/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func setupHandlers(userService user.UserService, authService auth.AuthService, productService product.ProductService) http.Handler {
	userHandler := &Handler.UserHandler{UserService: userService}
	authHandler := &Handler.AuthHandler{AuthService: authService}
	productHandler := &Handler.ProductHandler{ProductService: productService}

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.SecurityHeaders)

	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/refresh", authHandler.Refresh).Methods(http.MethodPost)
	r.HandleFunc("/products", productHandler.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/product", productHandler.GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/products/all", productHandler.GetAllProducts).Methods(http.MethodGet)

	return r
}
