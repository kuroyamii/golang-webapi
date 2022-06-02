package cafeServicePkg

import (
	"context"

	cafeDto "github.com/kuroyamii/golang-webapi/internal/cafe/dto"
)

type CafeService interface {
	GetAllFoodByType(ctx context.Context, food_type []string) (cafeDto.FoodsResponse, error)
	SearchFood(ctx context.Context, query string) (cafeDto.FoodsResponse, error)
	GetSeatData(ctx context.Context) (cafeDto.SeatsResponse, error)
	GetWaiterData(ctx context.Context) (cafeDto.WaitersResponse, error)
	// GetAllFood(ctx context.Context) (cafeDto.FoodsResponse, error)
	GetSumPeople(ctx context.Context) (cafeDto.PeopleSummary, error)
	GetCustomersOrderData(ctx context.Context) (cafeDto.CustomerDetails, error)
	GetCustomerOrderByCustomerID(ctx context.Context, customerID uint64) (cafeDto.OrderData, error)
	PlaceOrder(ctx context.Context, customerName string, tableID int, foodID []int, waiterID int, amount []int, price int) (uint64, error)
	PayBill(ctx context.Context, customerID uint64) error
	GetCustomerByID(ctx context.Context, customerID uint64) (cafeDto.CustomerResponse, error)
	GetFoodByFoodID(ctx context.Context, foodID int) (cafeDto.FoodResponse, error)
	GetEstimatedIncome(ctx context.Context) (cafeDto.EstimatedIncomesResponse, error)
	GetFoodTypes(ctx context.Context) (cafeDto.FoodTypesResponse, error)
	GetFoodByTypeAndName(ctx context.Context, name string, foodType []string) (cafeDto.FoodsResponse, error)
}
