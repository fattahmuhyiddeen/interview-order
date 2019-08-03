package config

import "os"

var Port = os.Getenv("PORT")
var DatabaseURL = os.Getenv("DATABASE_URL")
var Env = os.Getenv("ENV")
const DateTimeFormat = "2006-01-02T15:04:05"
