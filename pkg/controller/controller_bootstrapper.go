package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	cafeControllerPkg "github.com/kuroyamii/golang-webapi/internal/cafe/controller"
	cafeRepositoryPkg "github.com/kuroyamii/golang-webapi/internal/cafe/repository/impl"
	cafeServicePkg "github.com/kuroyamii/golang-webapi/internal/cafe/service/impl"
	pingController "github.com/kuroyamii/golang-webapi/internal/ping/controller"
	pingServicePkg "github.com/kuroyamii/golang-webapi/internal/ping/service"
	"github.com/kuroyamii/golang-webapi/pkg/middleware"

	authControllerPkg "github.com/kuroyamii/golang-webapi/internal/auth/controller"
	authServicePkg "github.com/kuroyamii/golang-webapi/internal/auth/service"
)

func InitializeController(router *mux.Router, db *sql.DB) {
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

	cafeRepository := cafeRepositoryPkg.ProvideCafeRepository(db)
	cafeService := cafeServicePkg.ProvideCafeService(cafeRepository, db)
	cafeController := cafeControllerPkg.ProvideController(db, webrouter, cafeService)
	cafeController.InitializeEndpoints()

}
