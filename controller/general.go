package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	model "github.com/fattahmuhyiddeen/interview-order/model"

	"github.com/labstack/echo"
)

//Timestamp is used to return timestamp and can also use to check whether system is running or not :P
func Timestamp(c echo.Context) (err error) {
	var results map[string]interface{}
	json.Unmarshal([]byte(`{ "human" : "`+time.Now().Format(time.RFC850)+`" , "miliseconds" : `+strconv.FormatInt(time.Now().UnixNano()/1000000, 10)+` }`), &results)
	return c.JSON(http.StatusOK, results)
}

func HomePage(c echo.Context) (err error) {
	var results map[string]interface{}
	json.Unmarshal([]byte(`{ "version" : "0.1.3"}`), &results)
	return c.JSON(http.StatusOK, results)
}

func CheckDB(c echo.Context) (err error) {
	var results map[string]interface{}
	var status = "failed"
	if model.TestConnection() {
		status = "success"
	}
	json.Unmarshal([]byte(`{ "connection to DB" : "`+status+`"}`), &results)
	return c.JSON(http.StatusOK, results)
}
