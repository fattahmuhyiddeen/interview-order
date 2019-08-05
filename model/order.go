package model

import (
	"log"
	"strconv"
)

//Order is a model
type Order struct {
	ID                   int    `json:"id" form:"id"`
	UserID               int    `json:"user_id" form:"user_id"`
	State                string `json:"state" form:"state"`
	ItemName             string `json:"item_name" form:"item_name"`
	Price                int    `json:"price" form:"price"`
	FrequencyUpdateOrder int    `json:"frequency_update_order" form:"frequency_update_order"`
	DeletedAt            string `json:"deleted_at" form:"deleted_at"`
	CreatedAt            string `json:"created_at" form:"created_at"`
	UpdatedAt            string `json:"updated_at" form:"updated_at"`
}

const orderTable = "orders"
const orderFields = "user_id, state, item_name, price, frequency_update_order, deleted_at, created_at, updated_at"

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
		order.CreatedAt,
		order.UpdatedAt,
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
			"', updated_at='" + DateTimeNow() +
			"' WHERE id=" + strconv.Itoa(order.ID) + " AND deleted_at IS NULL")

	if err != nil {
		log.Println(err)
	}
}

//ReadOrder is to get order by id
func ReadOrder(id int) (order Order) {
	connectDB()
	defer disconnectDB()

	db.QueryRow("SELECT id, "+orderFields+" FROM "+orderTable+" WHERE id=$1 AND deleted_at IS NULL", id).Scan(
		&order.ID,
		&order.UserID,
		&order.State,
		&order.ItemName,
		&order.Price,
		&order.FrequencyUpdateOrder,
		&order.DeletedAt,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	return
}

//DeleteOrder by id
func DeleteOrder(id int) {
	connectDB()
	defer disconnectDB()

	db.QueryRow("UPDATE " + orderTable + " SET deleted_at='" + DateTimeNow() + "' WHERE id=" + strconv.Itoa(id))
}

//ReadOrders return list of orders
func ReadOrders() (orders []Order) {
	connectDB()
	defer disconnectDB()

	rows, err := db.Query("SELECT id, " + orderFields + " FROM " + orderTable + " ORDER BY created_at")
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
				&tempOrder.DeletedAt,
				&tempOrder.CreatedAt,
				&tempOrder.UpdatedAt,
			)
			orders = append(orders, *tempOrder)
		}
	}
	return
}
