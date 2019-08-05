package config

import "os"

//Port is port number of this app is using
var Port = os.Getenv("PORT")

//DatabaseURL is postgresSQL URL
var DatabaseURL = os.Getenv("DATABASE_URL")

//Env is environment
var Env = os.Getenv("ENV")

//DateTimeFormat is date time format used accross the app
const DateTimeFormat = "2006-01-02T15:04:05"

//PaymentURL is URL to payment microservice
const PaymentURL = "https://interview-payment.herokuapp.com"
