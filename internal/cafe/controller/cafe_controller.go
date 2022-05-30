package cafeControllerPkg

import (
	"database/sql"
	"errors"

	// "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	cafeDto "github.com/kuroyamii/golang-webapi/internal/cafe/dto"
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

// func (cc *CafeController) handleGetFoodByType(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	food_type := vars["foodType"]
// 	foods, err := cc.cs.GetAllFoodByType(r.Context(), food_type)
// 	if err != nil {
// 		response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("Error", err.Error())).ToJSON(w)
// 		return
// 	}
// 	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foods).ToJSON(w)
// 	return
// }

func (cc *CafeController) handleGetFoodByQuery(w http.ResponseWriter, r *http.Request) {
	foodTypeRequest := new(cafeDto.FoodTypeRequestBody)
	err := foodTypeRequest.FromJSON(r.Body)
	if foodTypeRequest == nil {
		err = errors.New("forbidden query")
		response.NewErrorResponse(http.StatusForbidden, http.StatusText(http.StatusForbidden), response.NewErrorResponseValue("Error", err.Error())).ToJSON(w)
		return
	}
	if foodTypeRequest.FoodName != "" && len(foodTypeRequest.FoodType) == 0 {
		foods, err := cc.cs.SearchFood(r.Context(), foodTypeRequest.FoodName)
		if err != nil {
			if err.Error() == "no data found" {
				response.NewErrorResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), response.NewErrorResponseValue("Error", err.Error())).ToJSON(w)
				return
			}
			response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("Error", err.Error())).ToJSON(w)
			return
		}
		response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foods).ToJSON(w)
		return
	}
	if len(foodTypeRequest.FoodType) != 0 && foodTypeRequest.FoodName == "" {
		foods, err := cc.cs.GetAllFoodByType(r.Context(), foodTypeRequest.FoodType)
		if err != nil {
			response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("Error", err.Error())).ToJSON(w)
			return
		}
		response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foods).ToJSON(w)
		return
	}
	if foodTypeRequest.FoodName != "" && len(foodTypeRequest.FoodType) != 0 {
		foods, err := cc.cs.GetFoodByTypeAndName(r.Context(), foodTypeRequest.FoodName, foodTypeRequest.FoodType)
		if err != nil {
			response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("Error", err.Error())).ToJSON(w)
			return
		}
		response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foods).ToJSON(w)
		return
	}
}

func (cc *CafeController) handleGetSeats(w http.ResponseWriter, r *http.Request) {
	seats, err := cc.cs.GetSeatData(r.Context())
	if err != nil {
		response.NewErrorResponse(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, seats).ToJSON(w)
	return
}
func (cc *CafeController) handleGetWaiter(w http.ResponseWriter, r *http.Request) {
	waiters, err := cc.cs.GetWaiterData(r.Context())
	if err != nil {
		if err.Error() == "no data found" {
			response.NewErrorResponse(http.StatusNotFound,
				http.StatusText(http.StatusNotFound),
				response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
			return
		}
		response.NewErrorResponse(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, waiters).ToJSON(w)
	return
}

func (cc *CafeController) handleSumPeople(w http.ResponseWriter, r *http.Request) {
	sum, err := cc.cs.GetSumPeople(r.Context())
	if err != nil {
		if err.Error() == "no data found" {
			response.NewErrorResponse(http.StatusNotFound,
				http.StatusText(http.StatusNotFound),
				response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
			return
		}
		response.NewErrorResponse(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, sum).ToJSON(w)
	return
}

func (cc *CafeController) handleGetCustomersDetails(w http.ResponseWriter, r *http.Request) {
	data, err := cc.cs.GetCustomersOrderData(r.Context())
	if err != nil {
		if err.Error() == "no data found" {
			response.NewErrorResponse(http.StatusNotFound,
				http.StatusText(http.StatusNotFound),
				response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
			return
		}
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, data).ToJSON(w)
	return
}
func (cc *CafeController) handleOrderByCustomerID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	customerIDstring := query.Get("customerID")
	customerID, err := strconv.ParseUint(customerIDstring, 10, 64)
	data, err := cc.cs.GetCustomerOrderByCustomerID(r.Context(), customerID)
	if err != nil {
		if err.Error() == "no data found" {
			response.NewErrorResponse(http.StatusNotFound,
				http.StatusText(http.StatusNotFound),
				response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
			return
		}
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, data).ToJSON(w)
	return
}

func (cc *CafeController) handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
	orderRequest := new(cafeDto.OrderRequestBody)
	err := orderRequest.FromJSON(r.Body)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	err = cc.cs.PlaceOrder(r.Context(), orderRequest.CustomerName, orderRequest.TableID, orderRequest.FoodID)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, nil).ToJSON(w)
	return
}

func (cc *CafeController) handlePayBill(w http.ResponseWriter, r *http.Request) {
	cp := new(cafeDto.CustomerPay)
	err := cp.FromJSON(r.Body)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	err = cc.cs.PayBill(r.Context(), cp.CustomerID)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, nil).ToJSON(w)
	return
}

func (cc *CafeController) handleGetSingleCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customerID"]
	id, err := strconv.ParseUint(customerID, 10, 64)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	customer, err := cc.cs.GetCustomerByID(r.Context(), id)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, customer).ToJSON(w)
	return
}

func (cc *CafeController) handleFoodByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	foodID := vars["foodID"]
	id, err := strconv.ParseInt(foodID, 10, 64)
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	food, err := cc.cs.GetFoodByFoodID(r.Context(), int(id))

	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, food).ToJSON(w)
	return
}

func (cc *CafeController) handleEstimatedIncomes(w http.ResponseWriter, r *http.Request) {
	ei, err := cc.cs.GetEstimatedIncome(r.Context())
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, ei).ToJSON(w)
	return

}

func (cc *CafeController) handleGetTypes(w http.ResponseWriter, r *http.Request) {
	foodType, err := cc.cs.GetFoodTypes(r.Context())
	if err != nil {
		response.NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), response.NewErrorResponseValue("error", err.Error())).ToJSON(w)
		return
	}
	response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, foodType).ToJSON(w)
	return
}

func (cc *CafeController) InitializeEndpoints() {
	// cc.router.HandleFunc(global.API_GET_FOOD_BY_TYPE, cc.handleGetFoodByType).Methods(http.MethodGet)
	cc.router.HandleFunc(global.API_GET_FOOD_BY_QUERY, cc.handleGetFoodByQuery).Methods(http.MethodPost, http.MethodOptions) //approved
	cc.router.HandleFunc(global.API_GET_SEATS, cc.handleGetSeats).Methods(http.MethodGet, http.MethodOptions)
	cc.router.HandleFunc(global.API_GET_WAITERS, cc.handleGetWaiter).Methods(http.MethodGet, http.MethodOptions)
	cc.router.HandleFunc(global.API_GET_SUM_PEOPLE, cc.handleSumPeople).Methods(http.MethodGet, http.MethodOptions)
	cc.router.HandleFunc(global.API_GET_DETAILS, cc.handleGetCustomersDetails).Methods(http.MethodGet, http.MethodOptions)
	cc.router.HandleFunc(global.API_GET_DETAIL_BY_CUSTOMER_ID, cc.handleOrderByCustomerID).Methods(http.MethodGet, http.MethodOptions)
	cc.router.HandleFunc(global.API_POST_ORDER, cc.handlePlaceOrder).Methods(http.MethodPost, http.MethodOptions)
	cc.router.HandleFunc(global.API_POST_PAYBILL, cc.handlePayBill).Methods(http.MethodPost, http.MethodOptions)
	cc.router.HandleFunc(global.API_GET_CUSTOMER_BY_ID, cc.handleGetSingleCustomer).Methods(http.MethodGet, http.MethodOptions)
	cc.router.HandleFunc(global.API_GET_FOOD_BY_ID, cc.handleFoodByID).Methods(http.MethodGet, http.MethodOptions) //approved
	cc.router.HandleFunc(global.API_GET_TYPES, cc.handleGetTypes).Methods(http.MethodGet, http.MethodOptions)      //approved
}
