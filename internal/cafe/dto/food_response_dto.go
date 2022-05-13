package cafeDto

import "time"

// ---------------------------------------------------------------------------------->
//This Block of code is for defining default struct for data transfer
type FoodResponse struct {
	FoodID      int    `json:"foodID"`
	Name        string `json:"name"`
	ImagePath   string `json:"imagePath"`
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
	OrderID    uint64    `json:"orderID"`
	CustomerID uint64    `json:"customerID"`
	OrderedAt  time.Time `json:"orderedAt"`
	WaiterID   int       `json:"waiterID"`
}

type OrderDetailResponse struct {
	DetailID uint64 `json:"detailID"`
	OrderID  uint64 `json:"orderID"`
	FoodID   int    `json:"foodID"`
}

type OrderDetailsResponse []*OrderDetailsResponse

type FoodTypeResponse struct {
	FoodTypeID int    `json:"foodTypeID"`
	TypeName   string `json:"typeName"`
}

type FoodTypesResponse []*FoodTypeResponse

type RecordResponse struct {
	RecordID uint64 `json:"recordID"`
	FoodID   int    `json:"foodID"`
	Amount   int    `json:"amount"`
}

type RecordsResponse []*RecordResponse

type WaiterResponse struct {
	WaiterID int    `json:"waiterID"`
	Name     string `json:"name"`
}

type WaitersResponse []*WaiterResponse

type PeopleSummary struct {
	Sum int `json:"totalPeople"`
}

// ---------------------------------------------------------------------------------->
//This Block is for detailing customer's order
type OrderDetailsData struct {
	DetailID uint64    `json:"detailID"`
	FoodData *FoodData `json:"foodData"`
}

type OrderDetailsDatas []*OrderDetailsData

type FoodData struct {
	FoodID      int    `json:"foodID"`
	Name        string `json:"name"`
	ImagePath   string `json:"imagePath"`
	Price       int    `json:"price"`
	FoodType    string `json:"foodType"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

type OrderData struct {
	OrderID      uint64             `json:"orderID"`
	OrderedAt    time.Time          `json:"orderedAt"`
	WaiterData   *WaiterResponse    `json:"waiterData"`
	OrderDetails *OrderDetailsDatas `json:"orderDetails"`
}

type CustomerDetail struct {
	CustomerData *CustomerResponse `json:"customerData"`
	OrderData    *OrderData        `json:"orderData"`
}

type CustomerDetails []*CustomerDetail
