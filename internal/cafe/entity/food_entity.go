package cafeEntity

type Food struct {
	FoodID      int    `db:"food_id"`
	Name        string `db:"name"`
	ImagePath   string `db:"image_path"`
	Price       int    `db:"price"`
	FoodType    string `db:"type_name"`
	Description string `db:"description"`
	Stock       int    `db:"stock"`
}

type Foods []*Food

type FoodType struct {
	FoodTypeID int    `db:"food_type_id"`
	FoodType   string `db:"food_type"`
}

type FoodTypes []*FoodType
