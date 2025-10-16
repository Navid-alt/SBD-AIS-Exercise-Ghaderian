package repository

import "ordersystem/model"

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Hot Chocolate", Price: 2.50, Description: "Sweet chocolate drink"},
		{ID: 2, Name: "Water", Price: 1.00, Description: "Still mineral water"},
		{ID: 3, Name: "Iced Coffee", Price: 2.80, Description: "Cold and energizing"},
	}
	orders := []model.Order{
		{DrinkID: 1, Amount: 1, CreatedAt: "2025-10-10T12:15:00Z"},
		{DrinkID: 5, Amount: 2, CreatedAt: "2025-10-10T12:30:00Z"},
		{DrinkID: 2, Amount: 2, CreatedAt: "2025-10-10T13:00:00Z"},
	}
	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += uint64(order.Amount)
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
