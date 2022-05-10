package cafeRepositoryPkg

import (
	"context"

	cafeEntity "github.com/kuroyamii/golang-webapi/internal/cafe/entity"
)

type CafeRepository interface {
	GetAllFoodByType(ctx context.Context, food_type string) (cafeEntity.Foods, error)
}
