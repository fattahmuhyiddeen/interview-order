package service

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fattahmuhyiddeen/interview-order/config"
	"github.com/fattahmuhyiddeen/interview-order/model"
)

//CheckPayment is to status from payment microservice
func CheckPayment(order *model.Order) {
	// sleep to simulate delay in processing
	time.Sleep(15 * time.Second)
	response, err := http.Get(config.PaymentURL + "/payment_status")
	if err == nil {
		defer response.Body.Close()
		body, bodyErr := ioutil.ReadAll(response.Body)
		if bodyErr == nil {
			cleanResponse := string(string(body)[0])
			if order.State != "cancelled" {
				if cleanResponse == "0" {
					order.State = "declined"
				} else {
					order.State = "delivered"
				}
				model.UpdateAfterPayment(order.ID, order.State)
			}
		}
	}
}
