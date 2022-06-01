package cafeServicePkg

import (
	"context"
	"database/sql"
	"errors"
	"log"

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

func (cs cafeServiceImpl) SearchFood(ctx context.Context, query string) (cafeDto.FoodsResponse, error) {
	foods, err := cs.cr.SearchFood(ctx, query)
	if err != nil {

		return nil, err
	}

	if len(foods) < 1 {

		return nil, errors.New("no data found")
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

func (cs cafeServiceImpl) GetAllFoodByType(ctx context.Context, food_type []string) (cafeDto.FoodsResponse, error) {
	foods, err := cs.cr.GetAllFoodByType(ctx, food_type)
	// log.Println(foods)
	if err != nil {
		return nil, err
	}
	if len(foods) < 1 {
		return nil, errors.New("no data found")
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

func (cs cafeServiceImpl) GetSeatData(ctx context.Context) (cafeDto.SeatsResponse, error) {
	seats, err := cs.cr.GetSeats(ctx)
	if err != nil {
		return nil, err
	}
	seatsResponse := cafeDto.SeatsResponse{}
	for _, seat := range seats {
		var item cafeDto.SeatResponse
		item.TableID = seat.TableID
		item.Status = seat.Status
		seatsResponse = append(seatsResponse, &item)
	}
	return seatsResponse, nil
}

func (cs cafeServiceImpl) GetWaiterData(ctx context.Context) (cafeDto.WaitersResponse, error) {
	waiters, err := cs.cr.GetAllWaiter(ctx)
	if err != nil {
		return nil, err
	}
	if len(waiters) < 1 {
		return nil, errors.New("no data found")
	}
	waitersResponse := cafeDto.WaitersResponse{}
	for _, waiter := range waiters {
		var item cafeDto.WaiterResponse
		item.WaiterID = waiter.WaiterID
		item.Name = waiter.Name
		waitersResponse = append(waitersResponse, &item)

	}
	return waitersResponse, nil
}

func (cs cafeServiceImpl) GetSumPeople(ctx context.Context) (cafeDto.PeopleSummary, error) {
	sum, err := cs.cr.GetSumPeople(ctx)
	if err != nil {
		return cafeDto.PeopleSummary{}, err
	}
	summary := cafeDto.PeopleSummary{
		Sum: sum,
	}
	return summary, nil
}

func (cs cafeServiceImpl) GetCustomersOrderData(ctx context.Context) (cafeDto.CustomerDetails, error) {
	customers, err := cs.cr.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}
	// make empty struct array
	var customerDetails cafeDto.CustomerDetails
	for _, item := range customers {
		order, err := cs.cr.GetOrderByCustomerID(ctx, item.CustomerID)
		if err != nil {
			return nil, err
		}
		waiter, err := cs.cr.GetWaiterByWaiterID(ctx, order.WaiterID)
		if err != nil {
			return nil, err
		}
		od, err := cs.cr.GetOrderDetailsByOrderID(ctx, order.OrderID)
		if err != nil {
			return nil, err
		}
		var orderdetails cafeDto.OrderDetailsDatas
		for _, ods := range od {
			food, err := cs.cr.GetFoodByFoodID(ctx, ods.FoodID)
			if err != nil {
				return nil, err
			}
			od := &cafeDto.OrderDetailsData{
				DetailID: ods.DetailID,
				FoodData: &cafeDto.FoodData{
					FoodID:      food.FoodID,
					Name:        food.Name,
					ImagePath:   food.ImagePath,
					Price:       food.Price,
					FoodType:    food.FoodType,
					Description: food.Description,
					Stock:       food.Stock,
				},
			}
			orderdetails = append(orderdetails, od)
		}
		if err != nil {
			return nil, err
		}
		cust := &cafeDto.CustomerDetail{
			CustomerData: &cafeDto.CustomerResponse{
				CustomerID: item.CustomerID,
				Name:       item.Name,
				TableID:    item.TableID,
			},
			OrderData: &cafeDto.OrderData{
				OrderID:   order.OrderID,
				OrderedAt: order.OrderedAt,
				WaiterData: &cafeDto.WaiterResponse{
					WaiterID: waiter.WaiterID,
					Name:     waiter.Name,
				},
				OrderDetails: &orderdetails,
			},
		}
		customerDetails = append(customerDetails, cust)
	}
	return customerDetails, nil
}

func (cs cafeServiceImpl) GetCustomerOrderByCustomerID(ctx context.Context, customerID uint64) (cafeDto.OrderData, error) {
	order, err := cs.cr.GetOrderByCustomerID(ctx, customerID)
	if err != nil {
		return cafeDto.OrderData{}, nil
	}
	waiter, err := cs.cr.GetWaiterByWaiterID(ctx, order.WaiterID)
	if err != nil {
		return cafeDto.OrderData{}, err
	}
	od, err := cs.cr.GetOrderDetailsByOrderID(ctx, order.OrderID)
	if err != nil {
		return cafeDto.OrderData{}, err
	}
	var orderdetails cafeDto.OrderDetailsDatas
	for _, ods := range od {
		food, err := cs.cr.GetFoodByFoodID(ctx, ods.FoodID)
		if err != nil {
			return cafeDto.OrderData{}, err
		}
		od := &cafeDto.OrderDetailsData{
			DetailID: ods.DetailID,
			FoodData: &cafeDto.FoodData{
				FoodID:      food.FoodID,
				Name:        food.Name,
				ImagePath:   food.ImagePath,
				Price:       food.Price,
				FoodType:    food.FoodType,
				Description: food.Description,
				Stock:       food.Stock,
			},
		}
		orderdetails = append(orderdetails, od)
	}
	orderResponse := cafeDto.OrderData{
		OrderID:   order.OrderID,
		OrderedAt: order.OrderedAt,
		WaiterData: &cafeDto.WaiterResponse{
			WaiterID: waiter.WaiterID,
			Name:     waiter.Name,
		},
		OrderDetails: &orderdetails,
	}
	return orderResponse, nil
}

func (cs cafeServiceImpl) PlaceOrder(ctx context.Context, customerName string, tableID int, foodID []int, waiterID int, amount []int) (uint64, error) {
	err := cs.cr.ReserveTable(ctx, tableID)
	if err != nil {
		return 0, err
	}
	customerID, err := cs.cr.InsertCustomer(ctx, customerName, tableID)
	if err != nil {
		return 0, err
	}
	orderID, err := cs.cr.InsertOrder(ctx, customerID, waiterID)
	if err != nil {
		return 0, err
	}
	log.Println(orderID)
	err = cs.cr.InsertOrderDetails(ctx, orderID, foodID, amount)
	if err != nil {
		return 0, err
	}
	return customerID, nil
}

func (cs cafeServiceImpl) PayBill(ctx context.Context, customerID uint64) error {
	logs, err := cs.cr.TransferToLog(ctx, customerID)
	if err != nil {
		return err
	}
	tableID := logs[0].TableID
	err = cs.cr.RemoveCustomer(ctx, customerID, tableID)
	if err != nil {
		return err
	}
	return nil
}

func (cs cafeServiceImpl) GetCustomerByID(ctx context.Context, customerID uint64) (cafeDto.CustomerResponse, error) {
	cust, err := cs.cr.GetCustomerByCustomerID(ctx, customerID)
	if err != nil {
		return cafeDto.CustomerResponse{}, nil
	}
	customer := cafeDto.CustomerResponse{
		Name:       cust.Name,
		TableID:    cust.TableID,
		CustomerID: customerID,
	}
	return customer, nil
}
func (cs cafeServiceImpl) GetFoodByFoodID(ctx context.Context, foodID int) (cafeDto.FoodResponse, error) {
	food, err := cs.cr.GetFoodByFoodID(ctx, foodID)

	if err != nil {
		return cafeDto.FoodResponse{}, err
	}
	foodResponse := cafeDto.FoodResponse{
		FoodID:      food.FoodID,
		Name:        food.Name,
		ImagePath:   food.ImagePath,
		Price:       food.Price,
		FoodType:    food.FoodType,
		Description: food.Description,
		Stock:       food.Stock,
	}
	return foodResponse, nil
}

func (cs cafeServiceImpl) GetEstimatedIncome(ctx context.Context) (cafeDto.EstimatedIncomesResponse, error) {
	ei, err := cs.cr.GetEstimatedIncome(ctx)
	if err != nil {
		return nil, err
	}
	var eiResponse cafeDto.EstimatedIncomesResponse
	for _, item := range ei {
		i := cafeDto.EstimatedIncomeResponse{
			RecordID:        item.RecordID,
			FoodName:        item.FoodName,
			EstimatedIncome: item.EstimatedIncome,
		}
		eiResponse = append(eiResponse, &i)
	}
	return eiResponse, nil
}

func (cs cafeServiceImpl) GetFoodTypes(ctx context.Context) (cafeDto.FoodTypesResponse, error) {
	foodTypes, err := cs.cr.GetFoodType(ctx)
	if err != nil {
		return nil, err
	}
	var foodTypesResponse cafeDto.FoodTypesResponse
	for _, item := range foodTypes {
		foodType := cafeDto.FoodTypeResponse{
			TypeName: item.FoodType,
		}
		foodTypesResponse = append(foodTypesResponse, &foodType)
	}
	return foodTypesResponse, nil
}

func (cs cafeServiceImpl) GetFoodByTypeAndName(ctx context.Context, name string, foodType []string) (cafeDto.FoodsResponse, error) {
	foods, err := cs.cr.GetByTypeAndName(ctx, name, foodType)
	if err != nil {
		return nil, err
	}
	var foodsResponse cafeDto.FoodsResponse
	for _, item := range foods {
		food := cafeDto.FoodResponse{
			FoodID:      item.FoodID,
			Name:        item.Name,
			ImagePath:   item.ImagePath,
			Price:       item.Price,
			Description: item.Description,
			Stock:       item.Stock,
			FoodType:    item.FoodType,
		}
		foodsResponse = append(foodsResponse, &food)
	}
	return foodsResponse, nil
}
