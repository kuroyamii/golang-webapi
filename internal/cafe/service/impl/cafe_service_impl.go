package cafeServicePkg

import (
	"context"
	"database/sql"

	cafeDto "github.com/kuroyamii/golang-webapi/internal/cafe/dto"
	CafeRepository "github.com/kuroyamii/golang-webapi/internal/cafe/repository/api"
)

type cafeServiceImpl struct {
	cr CafeRepository.CafeRepository
	db *sql.DB
}

func ProvideCafeService(cr CafeRepository.CafeRepository, db *sql.DB) *cafeServiceImpl {
	return &cafeServiceImpl{
		db: db,
		cr: cr,
	}
}

func (cs cafeServiceImpl) SearchFood(ctx context.Context, query string)(cafeDto.FoodsResponse, error) {
	foods, err := cs.cr.SearchFood(ctx, query)
	if err != nil{
		return nil, err
	}
	foodsResponse := cafeDto.FoodsResponse{}
	for _, food := range foods{
		var item cafeDto.FoodResponse
		item.FoodID = food.FoodID
		item.Name = food.Name
		item.ImagePath = food.ImagePath
		item.Price = food.Price
		item.Description = food.Description
		item.Stock = food.Stock
		item.FoodType = food.FoodType
		foodsResponse = append(foodsResponse, &item)
	}
	return foodsResponse,nil
}

func (cs cafeServiceImpl) GetAllFoodByType(ctx context.Context, food_type string) (cafeDto.FoodsResponse, error) {
	foods, err := cs.cr.GetAllFoodByType(ctx, food_type)
	if err != nil {
		return nil, err
	}
	foodsResponse := cafeDto.FoodsResponse{}
	for _, food := range foods {
		var item cafeDto.FoodResponse
		item.FoodID = food.FoodID
		item.Name = food.Name
		item.ImagePath = food.ImagePath
		item.Price = food.Price
		item.Description = food.Description
		item.Stock = food.Stock
		item.FoodType = food.FoodType
		foodsResponse = append(foodsResponse, &item)
	}
	return foodsResponse, nil
}
