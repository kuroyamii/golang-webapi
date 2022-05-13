package cafeDto

import (
	"encoding/json"
	"io"
)

type OrderRequestBody struct {
	CustomerName string `json:"customerName" validate:"required"`
	TableID      int    `json:"tableID" validate:"required"`
	FoodID       []int  `json:"foodID" validate:"required"`
}

func (or *OrderRequestBody) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(or)
}
