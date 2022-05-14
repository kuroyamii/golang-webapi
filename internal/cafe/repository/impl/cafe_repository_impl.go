package cafeRepositoryPkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	cafeEntity "github.com/kuroyamii/golang-webapi/internal/cafe/entity"
	cafeQuery "github.com/kuroyamii/golang-webapi/internal/cafe/query"
)

type cafeRepositoryImpl struct {
	DB *sql.DB
}

func ProvideCafeRepository(db *sql.DB) *cafeRepositoryImpl {
	return &cafeRepositoryImpl{DB: db}
}

func (cr cafeRepositoryImpl) SearchFood(ctx context.Context, query string) (cafeEntity.Foods, error) {
	query = fmt.Sprint("%", query, "%")
	result, err := cr.DB.Query(cafeQuery.SEARCH_FOODS_BY_QUERY)
	if err != nil {
		log.Printf("ERROR Querying data -> query: %v, error: %v\n", query, err.Error())
		return nil, err
	}
	foods := cafeEntity.Foods{}

	for result.Next() {
		var foodItem cafeEntity.Food
		err = result.Scan(&foodItem.FoodID, &foodItem.Name, &foodItem.ImagePath, &foodItem.Price, &foodItem.FoodType, &foodItem.Description, &foodItem.Stock)
		if err != nil {
			log.Printf("ERROR Scanning data -> query: %v, error: %v\n", query, err.Error())
			return nil, err
		}
		foods = append(foods, &foodItem)
	}
	return foods, nil
}

func (cr cafeRepositoryImpl) GetAllFoodByType(ctx context.Context, food_type string) (cafeEntity.Foods, error) {
	result, err := cr.DB.Query(cafeQuery.GET_ALL_FOOD_BY_TYPE, food_type)
	if err != nil {
		log.Printf("ERROR Querying data -> foodType: %v, error: %v\n", food_type, err.Error())
		return nil, err
	}
	foods := cafeEntity.Foods{}

	for result.Next() {
		var foodItem cafeEntity.Food
		err = result.Scan(&foodItem.FoodID, &foodItem.Name, &foodItem.ImagePath, &foodItem.Price, &foodItem.FoodType, &foodItem.Description, &foodItem.Stock)
		if err != nil {
			log.Printf("ERROR Scanning data -> foodType: %v, error: %v\n", food_type, err.Error())
			return nil, err
		}
		foods = append(foods, &foodItem)

	}

	return foods, nil
}

func (cr cafeRepositoryImpl) ReserveTable(ctx context.Context, tableID int) error {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.RESERVE_TABLE)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> tableID: %v, error: %v\n", tableID, err.Error())
		return err
	}
	var res sql.Result
	res, err = stmt.ExecContext(ctx, tableID)
	if err != nil {
		log.Printf("ERROR Executing Statement -> tableID: %v, error: %v\n", tableID, err.Error())
		return err
	}
	var row int64
	row, err = res.RowsAffected()
	log.Printf("INFO Database -> Rows affected: %v", row)
	return nil
}

func (cr cafeRepositoryImpl) UnreserveTable(ctx context.Context, tableID int) error {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.UNRESERVE_TABLE)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> tableID: %v, error: %v\n", tableID, err.Error())
		return err
	}
	var res sql.Result
	res, err = stmt.ExecContext(ctx, tableID)
	if err != nil {
		log.Printf("ERROR Executing Statement -> tableID: %v, error: %v\n", tableID, err.Error())
		return err
	}
	var row int64
	row, err = res.RowsAffected()
	log.Printf("INFO Database -> Rows affected: %v", row)
	return nil
}

func (cr cafeRepositoryImpl) InsertCustomer(ctx context.Context, name string, reserveTable int) (uint64, error) {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.INSERT_CUSTOMER)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> customer name: %v, error: %v\n", name, err.Error())
		return 0, err
	}
	var res sql.Result
	res, err = stmt.ExecContext(ctx, name, reserveTable)
	if err != nil {
		log.Printf("ERROR Executing Statement -> customer name: %v, error: %v\n", name, err.Error())
		return 0, err
	}
	var row int64
	row, err = res.RowsAffected()
	log.Printf("INFO Database -> Rows affected: %v", row)
	row, err = res.LastInsertId()
	num := uint64(row)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (cr cafeRepositoryImpl) RemoveCustomer(ctx context.Context, customerID uint64, reserveTable int) error {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.REMOVE_CUSTOMER)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> customer ID: %v, error: %v\n", customerID, err.Error())
		return err
	}
	_, err = stmt.ExecContext(ctx, customerID)
	if err != nil {
		log.Printf("ERROR Executing Statement -> customer ID: %v, error: %v\n", customerID, err.Error())
		return err
	}
	err = cr.UnreserveTable(ctx, reserveTable)
	if err != nil {
		return err
	}
	return nil
}

func (cr cafeRepositoryImpl) GetOrderByCustomerID(ctx context.Context, customerID uint64) (cafeEntity.Order, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_ORDER_BY_CUSTOMER_ID, customerID)
	if err != nil {
		log.Printf("ERROR Querying -> customer ID: %v, error: %v\n", customerID, err.Error())
		return cafeEntity.Order{}, err
	}
	var order cafeEntity.Order
	for rows.Next() {
		err = rows.Scan(&order.OrderID, &order.CustomerID, &order.OrderedAt, &order.WaiterID)
		if err != nil {
			log.Printf("ERROR Scanning -> customer ID: %v, error: %v\n", customerID, err.Error())
			return cafeEntity.Order{}, err
		}
	}
	return order, nil
}

func (cr cafeRepositoryImpl) GetOrderDetailsByOrderID(ctx context.Context, orderID uint64) (cafeEntity.OrderDetails, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_ORDER_DETAILS_BY_ORDER_ID, orderID)
	if err != nil {
		log.Printf("ERROR Querying -> order ID: %v, error: %v\n", orderID, err.Error())
		return nil, err
	}
	var orderDetails cafeEntity.OrderDetails
	for rows.Next() {
		var orderDetail cafeEntity.OrderDetail
		err = rows.Scan(&orderDetail.DetailID, &orderDetail.OrderID, &orderDetail.FoodID)
		if err != nil {
			log.Printf("ERROR Scanning -> order ID: %v, error: %v\n", orderID, err.Error())
			return nil, err
		}
		orderDetails = append(orderDetails, &orderDetail)
	}

	return orderDetails, nil
}

func (cr cafeRepositoryImpl) InsertOrder(ctx context.Context, customerID uint64, sumOfWaiter int) (uint64, error) {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.INSERT_ORDER)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> customer ID: %v, error: %v\n", customerID, err.Error())
		return 0, err
	}
	// var res sql.Result
	rand.Seed(time.Now().UnixNano())
	res, err := stmt.ExecContext(ctx, customerID, rand.Intn(sumOfWaiter-1)+1)
	if err != nil {
		log.Printf("ERROR Executing Statement -> customer ID: %v, error: %v\n", customerID, err.Error())
		return 0, err
	}
	row, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := uint64(row)
	return id, nil
}

func (cr cafeRepositoryImpl) InsertOrderDetails(ctx context.Context, orderID uint64, foodID []int) error {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.INSERT_ORDER_DETAIL)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> order ID: %v, food ID: %v, error: %v\n", orderID, foodID, err.Error())
		return err
	}

	// var res sql.Result
	for _, food := range foodID {
		_, err = stmt.ExecContext(ctx, orderID, food)
		if err != nil {
			log.Printf("ERROR Executing Statement -> order ID: %v, food ID: %v, error: %v\n", orderID, food, err.Error())
			return err
		}
	}
	return nil
}

func (cr cafeRepositoryImpl) InsertRecord(ctx context.Context, foodID int, amount int) error {
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.INSERT_RECORD)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> food ID: %v, error: %v\n", foodID, err.Error())
		return err
	}
	_, err = stmt.ExecContext(ctx, foodID, amount)
	if err != nil {
		log.Printf("ERROR Executing Statement -> food ID: %v, error: %v\n", foodID, err.Error())
		return err
	}
	return nil
}

func (cr cafeRepositoryImpl) GetAllWaiter(ctx context.Context) (cafeEntity.Waiters, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_ALL_WAITER)
	if err != nil {
		log.Printf("ERROR Querying -> error: %v\n", err.Error())
		return nil, err
	}
	var waiters cafeEntity.Waiters
	for rows.Next() {
		var waiter cafeEntity.Waiter
		err = rows.Scan(&waiter.WaiterID, &waiter.Name)
		if err != nil {
			log.Printf("ERROR Scanning Rows -> error: %v\n", err.Error())
			return nil, err
		}
		waiters = append(waiters, &waiter)
	}
	return waiters, nil
}

func (cr cafeRepositoryImpl) GetSeats(ctx context.Context) (cafeEntity.Seats, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_SEATS)
	if err != nil {
		log.Printf("ERROR Querying -> error: %v\n", err.Error())
		return nil, err
	}
	var seats cafeEntity.Seats
	for rows.Next() {
		var seat cafeEntity.Seat
		err = rows.Scan(&seat.TableID, &seat.Status)
		if err != nil {
			log.Printf("ERROR Scanning Rows -> error: %v\n", err.Error())
			return nil, err
		}
		seats = append(seats, &seat)
	}
	return seats, nil
}

func (cr cafeRepositoryImpl) GetSumPeople(ctx context.Context) (int, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_SUM_PEOPLE)
	if err != nil {
		log.Printf("ERROR Querying -> error: %v\n", err.Error())
		return 0, err
	}
	var sum int
	for rows.Next() {
		err = rows.Scan(&sum)
		if err != nil {
			log.Printf("ERROR Scanning Rows -> error: %v\n", err.Error())
			return 0, err
		}
	}

	return sum, nil
}

func (cr cafeRepositoryImpl) GetCustomers(ctx context.Context) (cafeEntity.Customers, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_ALL_CUSTOMER)
	if err != nil {
		log.Printf("ERROR Querying -> error: %v\n", err.Error())
		return nil, err
	}
	var customers cafeEntity.Customers
	for rows.Next() {
		var cust cafeEntity.Customer
		err = rows.Scan(&cust.CustomerID, &cust.Name, &cust.TableID)
		if err != nil {
			log.Printf("ERROR Scanning Rows -> error: %v\n", err.Error())
			return nil, err
		}
		customers = append(customers, &cust)
	}
	return customers, nil
}

func (cr cafeRepositoryImpl) GetWaiterByWaiterID(ctx context.Context, waiterID int) (cafeEntity.Waiter, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_WAITER_BY_WAITER_ID, waiterID)
	if err != nil {
		log.Printf("ERROR Querying -> waiterID: %v, error: %v\n", waiterID, err.Error())
		return cafeEntity.Waiter{}, err
	}
	var waiter cafeEntity.Waiter
	for rows.Next() {
		err = rows.Scan(&waiter.WaiterID, &waiter.Name)
		if err != nil {
			log.Printf("ERROR Scanning Rows ->  waiterID: %v, error: %v\n", waiterID, err.Error())
			return cafeEntity.Waiter{}, err
		}
	}
	return waiter, nil
}

func (cr cafeRepositoryImpl) GetFoodByFoodID(ctx context.Context, foodID int) (cafeEntity.Food, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_FOOD_BY_ID, foodID)
	if err != nil {
		log.Printf("ERROR Querying -> foodID: %v, error: %v\n", foodID, err.Error())
		return cafeEntity.Food{}, err
	}
	var food cafeEntity.Food
	for rows.Next() {
		err = rows.Scan(&food.FoodID, &food.Name, &food.ImagePath, &food.Price, &food.FoodType, &food.Description, &food.Stock)
		log.Printf("ERROR Scanning Rows ->  foodID: %v, error: %v\n", foodID, err.Error())
		return cafeEntity.Food{}, err
	}
	return food, nil
}

func (cr cafeRepositoryImpl) GetSumWaiter(ctx context.Context) (int, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_WAITER_SUM)
	if err != nil {
		log.Printf("ERROR Querying -> error: %v\n", err.Error())
		return 0, err
	}
	var sum int
	for rows.Next() {
		err = rows.Scan(&sum)
		if err != nil {
			log.Printf("ERROR Scanning Rows -> error: %v\n", err.Error())
		}
	}
	return sum, nil
}

func (cr cafeRepositoryImpl) TransferToLog(ctx context.Context, customerID uint64) (cafeEntity.Logs, error) {
	customer, err := cr.GetCustomerByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	order, err := cr.GetOrderByCustomerID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	od, err := cr.GetOrderDetailsByOrderID(ctx, order.OrderID)
	if err != nil {
		return nil, err
	}
	var logs cafeEntity.Logs
	for _, item := range od {
		var log cafeEntity.Log
		log = cafeEntity.Log{
			CustomerID:   customer.CustomerID,
			CustomerName: customer.Name,
			TableID:      customer.TableID,
			OrderID:      order.OrderID,
			WaiterID:     order.WaiterID,
			OrderedAt:    order.OrderedAt,
			DetailsID:    item.DetailID,
			FoodID:       item.FoodID,
		}
		logs = append(logs, &log)
	}
	stmt, err := cr.DB.PrepareContext(ctx, cafeQuery.INSERT_TO_LOG)
	if err != nil {
		log.Printf("ERROR Preparing Statement -> customerID: %v, error: %v\n", customerID, err.Error())
		return nil, err
	}
	var res sql.Result
	for _, item := range logs {
		res, err = stmt.ExecContext(ctx, item.CustomerID, item.CustomerName, item.TableID, item.OrderID, item.WaiterID, item.OrderedAt, item.DetailsID, item.FoodID)
		if err != nil {
			log.Printf("ERROR Executing Statement -> customerID: %v, error: %v\n", customerID, err.Error())
			return nil, err
		}
		log.Println(res.LastInsertId())
	}
	return logs, nil
}

func (cr cafeRepositoryImpl) GetCustomerByCustomerID(ctx context.Context, customerID uint64) (cafeEntity.Customer, error) {
	rows, err := cr.DB.QueryContext(ctx, cafeQuery.GET_CUSTOMER_BY_ID, customerID)
	if err != nil {
		log.Printf("ERROR Querying -> customerID: %v, error: %v\n\n", customerID, err.Error())
	}
	var cust cafeEntity.Customer
	for rows.Next() {
		err = rows.Scan(&cust.CustomerID, &cust.Name, &cust.TableID)
		if err != nil {
			log.Printf("ERROR Scanning Rows -> customerID: %v, error: %v\n", customerID, err.Error())
		}
	}
	return cust, nil
}
