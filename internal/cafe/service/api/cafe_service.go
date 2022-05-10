package cafeServicePkg

import (
	"context"

	cafeDto "github.com/kuroyamii/golang-webapi/internal/cafe/dto"
)

type CafeService interface {
	GetAllFoodByType(ctx context.Context, food_type string) (cafeDto.FoodsResponse, error)
}
