package routed

import (
	"database/sql"
	"fmt"
	"net/http"
)

func InitRoutes(db *sql.DB) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "OK")
	})
	http.HandleFunc("/order", handleAddOrder(db))
	http.HandleFunc("/pay", handleAddReceipt(db))
	http.HandleFunc("/click", handlerClick(db))
	http.HandleFunc("/delivered", handlerDelivered(db))
}
