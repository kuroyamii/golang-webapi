package cafeEntity

type Food struct {
	FoodID      int    `db:"food_id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	FoodType    string `db:"food_type"`
	Description string `db:"description"`
	Stock       int    `db:"stock"`
}

type Foods []*Food
