package cafeDto

import "time"

type FoodResponse struct {
	FoodID      int    `json:"foodID"`
	Name        string `json:"name"`
	ImagePath string `json:"imagePath"`
	Price       int    `json:"price"`
	FoodType    string `json:"foodType"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

type FoodsResponse []*FoodResponse

type SeatResponse struct {
	TableID int  `json:"tableID"`
	Status  bool `json:"status"`
}

type SeatsResponse []*SeatResponse

type CustomerResponse struct {
	CustomerID uint64 `json:"customerID"`
	Name       string `json:"name"`
	TableID    int    `json:"tableID"`
}
type CustomersResponse []*CustomerResponse

type OrderResponse struct {
	OrderID    uint64 `json:"orderID"`
	CustomerID uint64 `json:"customerID"`
	OrderedAt  time.Time `json:"orderedAt"`
}

type OrderDetailResponse struct{
	DetailID uint64 `json:"detailID"`
	OrderID uint64 `json:"orderID"`
	FoodID int `json:"foodID"`
}

type OrderDetailsResponse []*OrderDetailResponse

type FoodTypeResponse struct{
	FoodTypeID int `json:"foodTypeID"`
	TypeName string `json:"typeName"`
}

type FoodTypesResponse []*FoodTypeResponse


type RecordResponse struct{
	RecordID uint64 `json:"recordID"`
	FoodID int `json:"foodID"`
	Amount int `json:"amount"`
}

type RecordsResponse []*RecordResponse