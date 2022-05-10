package cafeRepositoryPkg

import (
	"context"
	"database/sql"
	"log"

	cafeEntity "github.com/kuroyamii/golang-webapi/internal/cafe/entity"
)

const (
	GET_ALL_FOOD_BY_TYPE = `SELECT f.food_id,f.name,f.image_path,f.price,t.type_name,f.description,f.stock
	FROM food f
	INNER JOIN food_type t ON t.food_type_id = f.food_type_id
	WHERE t.type_name = ?;`
)

type cafeRepositoryImpl struct {
	DB *sql.DB
}

func ProvideCafeRepository(db *sql.DB) *cafeRepositoryImpl {
	return &cafeRepositoryImpl{DB: db}
}

func (cr cafeRepositoryImpl) GetAllFoodByType(ctx context.Context, food_type string) (cafeEntity.Foods, error) {
	result, err := cr.DB.Query(GET_ALL_FOOD_BY_TYPE, food_type)
	if err != nil {
		log.Printf("ERROR Querying data -> foodType: %v, error: %v", food_type, err.Error())
		return nil, err
	}
	foods := cafeEntity.Foods{}

	for result.Next() {
		var foodItem cafeEntity.Food
		err = result.Scan(&foodItem.FoodID, &foodItem.Name,&foodItem.ImagePath, &foodItem.Price, &foodItem.FoodType, &foodItem.Description, &foodItem.Stock)
		if err != nil {
			log.Printf("ERROR Scanning data -> foodType: %v, error: %v", food_type, err.Error())
			return nil, err
		}
		foods = append(foods, &foodItem)
	}
	log.Println(foods)

	return foods, nil
}
