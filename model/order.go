package model

import (
	"log"
	"strconv"
	"time"

	config "github.com/fattahmuhyiddeen/interview-order/config"
)

//Order is a model
type Order struct {
	ID                   int       `json:"id" form:"id"`
	UserID               int       `json:"user_id" form:"user_id"`
	State                string    `json:"state" form:"state"`
	ItemName             string    `json:"item_name" form:"item_name"`
	Price                int       `json:"price" form:"price"`
	FrequencyUpdateOrder int       `json:"frequency_update_order" form:"frequency_update_order"`
	CreatedAt            time.Time `json:"created_at" form:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" form:"updated_at"`
}

const orderTable = "orders"
const orderFields = "user_id, state, item_name, price, frequency_update_order, created_at, updated_at"

//CreateOrder inserts a new order into orders table
func CreateOrder(order *Order) {
	connectDB()
	defer disconnectDB()

	order.CreatedAt = DateTimeNow()
	order.UpdatedAt = order.CreatedAt

	err := db.QueryRow(
		"INSERT INTO "+orderTable+" ("+orderFields+") VALUES ("+tablePlaceholder(orderFields)+") RETURNING id",
		order.UserID,
		order.State,
		order.ItemName,
		order.Price,
		order.FrequencyUpdateOrder,
		nil,
		order.CreatedAt.Format(config.DateTimeFormat),
		order.UpdatedAt.Format(config.DateTimeFormat),
	).Scan(&order.ID)

	if err != nil {
		log.Println(err)
	}
}

//UpdateOrder updates an order in orders table
func UpdateOrder(order *Order) {
	connectDB()
	defer disconnectDB()

	err := db.QueryRow(
		"UPDATE " + orderTable +
			" SET user_id='" + strconv.Itoa(order.UserID) +
			"', state='" + order.State +
			"', item_name='" + order.ItemName +
			"', price='" + strconv.Itoa(order.Price) +
			// "', frequency_update_order='" + strconv.Itoa(order.FrequencyUpdateOrder) +
			"', updated_at='" + DateTimeNow().Format(config.DateTimeFormat) +
			"' WHERE id=" + strconv.Itoa(order.ID))

	if err != nil {
		log.Println(err)
	}
}

//ReadOrder is to get order by id
func ReadOrder(id int) (order Order) {
	connectDB()
	defer disconnectDB()

	db.QueryRow("SELECT id, "+orderFields+" FROM "+orderTable+" WHERE id=$1", id).Scan(
		&order.ID,
		&order.UserID,
		&order.State,
		&order.ItemName,
		&order.Price,
		&order.FrequencyUpdateOrder,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	return
}

//UpdateAfterPayment is to either decline/deliver order after payment result
func UpdateAfterPayment(id int, state string) {
	connectDB()
	defer disconnectDB()
	db.QueryRow("UPDATE " + orderTable + " SET state='" + state + "' WHERE id=" + strconv.Itoa(id) + " AND state<>'cancelled'")
}

//ReadOrders return list of orders
func ReadOrders() (orders []Order) {
	connectDB()
	defer disconnectDB()

	rows, err := db.Query("SELECT id, " + orderFields + " FROM " + orderTable + " ORDER BY id")
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			tempOrder := new(Order)
			rows.Scan(
				&tempOrder.ID,
				&tempOrder.UserID,
				&tempOrder.State,
				&tempOrder.ItemName,
				&tempOrder.Price,
				&tempOrder.FrequencyUpdateOrder,
				&tempOrder.CreatedAt,
				&tempOrder.UpdatedAt,
			)
			orders = append(orders, *tempOrder)
		}
	}
	return
}
