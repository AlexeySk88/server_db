package routed

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func InitRoutes(db *sql.DB, r *chi.Mux) {
	r.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "OK")
	})
	r.Post("/order", handleAddOrder(db))
	r.Post("/pay", handleAddReceipt(db))
	r.Post("/click", handlerClick(db))
	r.Post("/delivered", handlerDelivered(db))
}
