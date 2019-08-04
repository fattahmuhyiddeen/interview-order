package controller

import (
	"net/http"

	model "github.com/fattahmuhyiddeen/interview-order/model"

	"github.com/labstack/echo"
)

//CreateOrder is to create new order
func CreateOrder(c echo.Context) (err error) {
	order := new(model.Order)
	c.Bind(order)
	model.CreateOrder(order)
	return c.JSON(http.StatusCreated, order)

}
