package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/idrpambudi/fita-appointment/bootstrap"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
