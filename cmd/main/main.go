package main

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kuroyamii/golang-webapi/pkg/controller"
	"github.com/kuroyamii/golang-webapi/pkg/database"
	"github.com/kuroyamii/golang-webapi/pkg/middleware"
	"github.com/kuroyamii/golang-webapi/pkg/server"
)

func getEnvVariables() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")
	envVariables["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	envVariables["DB_USERNAME"] = os.Getenv("DB_USERNAME")
	envVariables["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVariables["DB_NAME"] = os.Getenv("DB_NAME")
	envVariables["WHITELISTED_URLS"] = os.Getenv("WHITELISTED_URLS")

	return envVariables
}

func initializeGlobalRouter(envVariables map[string]string) *mux.Router {
	router := mux.NewRouter()

	arrayWhitelistedUrls := strings.Split(envVariables["WHITELISTED_URLS"], ",")
	whitelistedUrls := make(map[string]bool)

	for _, values := range arrayWhitelistedUrls {
		whitelistedUrls[values] = true
	}
	router.Use(middleware.ContentTypeMiddleware)
	router.Use(middleware.CorsMiddleware(whitelistedUrls))
	return router
}

func initializeDatabase(envVariables map[string]string) *sql.DB {
	return database.GetDatabase(
		envVariables["DB_ADDRESS"],
		envVariables["DB_USERNAME"],
		envVariables["DB_PASSWORD"],
		envVariables["DB_NAME"])
}

func main() {
	godotenv.Load()
	envVariables := getEnvVariables()
	router := initializeGlobalRouter(envVariables)
	db := initializeDatabase(envVariables)
	controller.InitializeController(router, db)

	server := server.ProvideServer(envVariables["SERVER_ADDRESS"], router)
	server.ListenAndServe()

}
