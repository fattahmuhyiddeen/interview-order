package service

import (
	"io/ioutil"
	"net/http"
	// "time"
	"github.com/fattahmuhyiddeen/interview-order/config"
	"github.com/fattahmuhyiddeen/interview-order/model"
)

//CheckPayment is to status from payment microservice
func CheckPayment(order *model.Order) {
	// sleep to simulate delay in processing
	// time.Sleep(15 * time.Second)
	response, err := http.Get(config.PaymentURL + "/payment_status")
	if err == nil {
		defer response.Body.Close()
		body, bodyErr := ioutil.ReadAll(response.Body)
		if bodyErr == nil {
			if order.State != "cancelled" {
				if string(body) == "0" {
					order.State = "declined"
				} else if string(body) == "1" {
					order.State = "delivered"
				}
				model.UpdateOrder(order)

			}
		}
	}
}
