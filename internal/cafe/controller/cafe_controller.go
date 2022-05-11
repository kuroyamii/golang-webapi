package cafeControllerPkg

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	cafeService "github.com/kuroyamii/golang-webapi/internal/cafe/service/api"
	"github.com/kuroyamii/golang-webapi/internal/global"
	"github.com/kuroyamii/golang-webapi/pkg/entity/response"
)

type CafeController struct {
	router *mux.Router
	db     *sql.DB
	cs     cafeService.CafeService
}

func ProvideController(db *sql.DB, r *mux.Router, cs cafeService.CafeService) *CafeController {

	return &CafeController{
		router: r,
		db:     db,
		cs:     cs,
	}
}

func (cc *CafeController) handleGetFoodByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	food_type := vars["foodType"]
	foods, err := cc.cs.GetAllFoodByType(r.Context(), food_type)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("Error", err.Error()))
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foods).ToJSON(w)
	return
}

func (cc *CafeController) handleGetFoodByQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	queryText := query.Get("foodSearch")
	foods, err := cc.cs.SearchFood(r.Context(), queryText)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("Error", err.Error()))
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foods).ToJSON(w)
	return
}

func (cc *CafeController) InitializeEndpoints() {
	cc.router.HandleFunc(global.API_GET_FOOD_BY_TYPE, cc.handleGetFoodByType).Methods(http.MethodGet)
	cc.router.HandleFunc(global.API_GET_FOOD_BY_QUERY, cc.handleGetFoodByQuery).Methods(http.MethodGet)
}
