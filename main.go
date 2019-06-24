package main

import (
	"fmt"
	"log"
	"net/http"
	"server_db/config"
	"server_db/db"
	"server_db/routed"
	"server_db/service"
)

func main() {
	config.Set()

	conSQL, errDB := db.Connect()
	if errDB != nil {
		fmt.Printf("failed to connect db: %v\n", errDB)
	}
	service.AddTables(conSQL)
	routed.InitRoutes(conSQL)

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
