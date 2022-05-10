package cafeDto

type FoodResponse struct {
	FoodID      int    `json:"foodID"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	FoodType    string `json:"foodType"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

type FoodsResponse []*FoodResponse
