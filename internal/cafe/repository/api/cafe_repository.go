package cafeRepositoryPkg

import (
	"context"

	cafeEntity "github.com/kuroyamii/golang-webapi/internal/cafe/entity"
)

type CafeRepository interface {
	GetAllFoodByType(ctx context.Context, food_type []string) (cafeEntity.Foods, error)
	SearchFood(ctx context.Context, query string) (cafeEntity.Foods, error)
	ReserveTable(ctx context.Context, tableID int) error
	UnreserveTable(ctx context.Context, tableID int) error
	InsertCustomer(ctx context.Context, name string, reserveTable int) (uint64, error)
	RemoveCustomer(ctx context.Context, customerID uint64, reserveTable int) error
	GetOrderByCustomerID(ctx context.Context, customerID uint64) (cafeEntity.Order, error)
	InsertOrder(ctx context.Context, customerID uint64, waiterID int, price int) (uint64, error)
	GetOrderDetailsByOrderID(ctx context.Context, orderID uint64) (cafeEntity.OrderDetails, error)
	InsertOrderDetails(ctx context.Context, orderID uint64, foodID []int, amount []int) error
	// InsertRecord(ctx context.Context, foodID int, amount int) error
	GetAllWaiter(ctx context.Context) (cafeEntity.Waiters, error)
	GetSeats(ctx context.Context) (cafeEntity.Seats, error)
	GetSumPeople(ctx context.Context) (int, error)
	GetCustomers(ctx context.Context) (cafeEntity.Customers, error)
	GetWaiterByWaiterID(ctx context.Context, waiterID int) (cafeEntity.Waiter, error)
	GetFoodByFoodID(ctx context.Context, foodID int) (cafeEntity.Food, error)
	GetSumWaiter(ctx context.Context) (int, error)
	TransferToLog(ctx context.Context, customerID uint64) (cafeEntity.Logs, error)
	GetCustomerByCustomerID(ctx context.Context, customerID uint64) (cafeEntity.Customer, error)
	GetEstimatedIncome(ctx context.Context) (cafeEntity.EstimatedIncomes, error)
	FilterFoodByPrice(ctx context.Context, min int, max int, mode string) (cafeEntity.Foods, error)
	GetFoodType(ctx context.Context) (cafeEntity.FoodTypes, error)
	GetByTypeAndName(ctx context.Context, name string, foodType []string) (cafeEntity.Foods, error)
	Restock(ctx context.Context) error
}
