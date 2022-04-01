package controller

import (
	"github.com/gorilla/mux"
	pingController "github.com/kuroyamii/golang-webapi/internal/ping/controller"
	pingServicePkg "github.com/kuroyamii/golang-webapi/internal/ping/service"
	"github.com/kuroyamii/golang-webapi/pkg/middleware"
)

func InitializeController(router *mux.Router) {
	router.Use(middleware.ErrorHandlingMiddleware)

	//init web router path
	webrouter := router.PathPrefix(API_PATH).Subrouter()

	//initialize ping controller
	pingService := pingServicePkg.ProvidePingService()
	pingController := pingController.ProvidePingController(webrouter, &pingService)
	pingController.InitializeEndPoint()

}
