package cafeEntity

import "time"

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
	FoodType   string `db:"type_name"`
}

type FoodTypes []*FoodType

type Customer struct {
	CustomerID uint64 `db:"customer_id"`
	Name       string `db:"name"`
	TableID    int    `db:"table_id"`
}

type Customers []*Customer

type Seat struct {
	TableID int  `db:"table_id"`
	Status  bool `db:"status"`
}

type Seats []*Seat

type Order struct {
	OrderID    uint64    `db:"order_id"`
	CustomerID uint64    `db:"customer_id"`
	OrderedAt  time.Time `db:"ordered_at"`
	WaiterID   int       `db:"waiter_id"`
}

type OrderDetail struct {
	DetailID uint64 `db:"details_id"`
	OrderID  uint64 `db:"order_id"`
	FoodID   int    `db:"food_id"`
}
type OrderDetails []*OrderDetail

type Record struct {
	RecordID uint64 `db:"record_id"`
	FoodID   int    `db:"food_id"`
	Amount   int    `db:"amount"`
}

type Records []*Record

type Waiter struct {
	WaiterID int    `db:"waiter_id"`
	Name     string `db:"name"`
}

type Waiters []*Waiter

type Log struct {
	CustomerID   uint64    `db:"customer_id"`
	CustomerName string    `db:"customer_name"`
	TableID      int       `db:"table_id"`
	OrderID      uint64    `db:"order_id"`
	WaiterID     int       `db:"waiter_id"`
	OrderedAt    time.Time `db:"ordered_at"`
	DetailsID    uint64    `db:"details_id"`
	FoodID       int       `db:"food_id"`
}

type Logs []*Log

type EstimatedIncome struct {
	RecordID        uint64 `db:"record_id"`
	FoodName        string `db:"food_name"`
	EstimatedIncome int    `db:"estimated_income"`
}
type EstimatedIncomes []*EstimatedIncome
