package cafeServicePkg

import (
	"context"

	cafeDto "github.com/kuroyamii/golang-webapi/internal/cafe/dto"
)

type CafeService interface {
	GetAllFoodByType(ctx context.Context, food_type string) (cafeDto.FoodsResponse, error)
	SearchFood(ctx context.Context, query string) (cafeDto.FoodsResponse, error)
	GetSeatData(ctx context.Context) (cafeDto.SeatsResponse, error)
	GetWaiterData(ctx context.Context) (cafeDto.WaitersResponse, error)
	// GetAllFood(ctx context.Context) (cafeDto.FoodsResponse, error)
	GetSumPeople(ctx context.Context) (cafeDto.PeopleSummary, error)
	GetCustomersOrderData(ctx context.Context) (cafeDto.CustomerDetails, error)
	GetCustomerOrderByCustomerID(ctx context.Context, customerID uint64) (cafeDto.OrderData, error)
}
