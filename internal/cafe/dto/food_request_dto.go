package cafeDto

import (
	"encoding/json"
	"io"
)

type OrderRequestBody struct {
	CustomerName string `json:"customerName" validate:"required"`
	TableID      int    `json:"tableID" validate:"required"`
	FoodID       []int  `json:"foodID" validate:"required"`
	WaiterID     int    `json:"waiterID" validate:"required"`
	Amount       []int  `json:"amount" validate:"required"`
}

type CustomerPay struct {
	CustomerID uint64 `json:"customerID"`
}

func (or *OrderRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(or)
}
func (cp *CustomerPay) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(cp)
}

type FoodTypeRequestBody struct {
	FoodType []string `json:"foodType"`
	FoodName string   `json:"foodName"`
}

func (ft *FoodTypeRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(ft)
}
