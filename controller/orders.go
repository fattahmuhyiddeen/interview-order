package controller

import (
	"net/http"

	model "github.com/fattahmuhyiddeen/interview-order/model"

	"github.com/labstack/echo"
)

//RequestOrder is request struct of order
type RequestOrder struct {
	model.Order
}

//CreateOrder is to create new order
func CreateOrder(c echo.Context) (err error) {
	order := new(model.Order)
	c.Bind(order)
	order.State = "confirmed"
	model.CreateOrder(order)
	return c.JSON(http.StatusCreated, order)
}

//CheckOrder is to get order by id
func CheckOrder(c echo.Context) (err error) {
	request := new(RequestOrder)
	c.Bind(request)
	order := model.ReadOrder(request.ID)
	return c.JSON(http.StatusOK, order)
}

//CancelOrder is to cancel order by id
func CancelOrder(c echo.Context) (err error) {
	request := new(RequestOrder)
	c.Bind(request)
	order := model.ReadOrder(request.ID)
	order.State = "cancelled"
	model.UpdateOrder(&order)
	// model.DeleteOrder(request.ID)
	return c.JSON(http.StatusOK, "order")
}

//ListOrders is to return list of orders
func ListOrders(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, model.ReadOrders())
}
