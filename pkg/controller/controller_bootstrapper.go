package controller

import (
	"github.com/gorilla/mux"
	pingController "github.com/kuroyamii/golang-webapi/internal/ping/controller"
	pingServicePkg "github.com/kuroyamii/golang-webapi/internal/ping/service"
	"github.com/kuroyamii/golang-webapi/pkg/middleware"

	authControllerPkg "github.com/kuroyamii/golang-webapi/internal/auth/controller"
	authServicePkg "github.com/kuroyamii/golang-webapi/internal/auth/service"
)

func InitializeController(router *mux.Router) {
	router.Use(middleware.ErrorHandlingMiddleware)

	//init web router path
	webrouter := router.PathPrefix(API_PATH).Subrouter()

	//initialize ping controller
	pingService := pingServicePkg.ProvidePingService()
	pingController := pingController.ProvidePingController(webrouter, &pingService)
	pingController.InitializeEndPoint()

	authService := authServicePkg.ProvideAuthService()
	authController := authControllerPkg.ProvideAuthController(webrouter, &authService)
	authController.InitializeEndPoint()

}
