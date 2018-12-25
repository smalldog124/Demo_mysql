package main

import (
	"Demo_mysql/route"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
)

const (
	username = "root"
	password = "sckshuhari"
	host     = "localhost"
	port     = "3360"
	database = "smalldogShop"
)

func main() {
	urlSql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err := sqlx.Connect("mysql", urlSql)
	if err != nil {
		log.Println(err)
		log.Fatal("Can not Connet Database")
	}

	router := mux.NewRouter()
	route.UserRouter(router, db)
	http.ListenAndServe(":3000", router)
}
