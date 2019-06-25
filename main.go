package main

import (
	"fmt"
	"log"
	"net/http"
	"server_db/config"
	"server_db/db"
	"server_db/routed"
	"server_db/service"

	"github.com/go-chi/chi"
)

func main() {
	config.Set()

	conSQL, errDB := db.Connect()
	if errDB != nil {
		fmt.Printf("failed to connect db: %v\n", errDB)
	}
	service.AddTables(conSQL)
	r := chi.NewRouter()
	routed.InitRoutes(conSQL, r)

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
