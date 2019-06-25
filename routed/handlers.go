package routed

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"server_db/service"
)

type ErrorRes struct {
	Res    string
	Report interface{}
}

func handleAddOrder(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		reqJSON := []byte(req.FormValue("order"))
		order := service.OrderValue{}
		errType := json.Unmarshal(reqJSON, &order)
		if errType != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: errType})
			res.Write(errorJSON)
			return
		}

		result, err := service.Order(db, order)
		if err != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: err})
			res.Write(errorJSON)
			return
		}
		resultJSON, _ := json.Marshal(result)
		res.Write(resultJSON)
	}
}

func handleAddReceipt(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		reqJSON := []byte(req.FormValue("receipt"))
		receipt := service.ReceiptValue{}
		errType := json.Unmarshal(reqJSON, &receipt)
		if errType != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: errType})
			res.Write(errorJSON)
			return
		}

		result, err := service.Pay(db, receipt)
		if err != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: fmt.Sprint("Error", err)})
			res.Write(errorJSON)
			return
		}
		resultJSON, _ := json.Marshal(result)
		res.Write(resultJSON)
	}
}

func handlerClick(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		reqJSON := []byte(req.FormValue("click"))
		click := service.ClickValue{}
		errType := json.Unmarshal(reqJSON, &click)
		if errType != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: errType})
			res.Write(errorJSON)
			return
		}

		result, err := service.Click(db, click)
		if err != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: fmt.Sprint("Error", err)})
			res.Write(errorJSON)
			return
		}
		resultJSON, _ := json.Marshal(result)
		res.Write(resultJSON)
	}
}

func handlerDelivered(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "method not supported", http.StatusMethodNotAllowed)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		reqJSON := []byte(req.FormValue("delivered"))
		deliv := service.DeliveryValue{}
		errType := json.Unmarshal(reqJSON, &deliv)
		if errType != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: errType})
			res.Write(errorJSON)
			return
		}

		result, err := service.Delivered(db, deliv)
		if err != nil {
			errorJSON, _ := json.Marshal(ErrorRes{Res: "fail", Report: err})
			res.Write(errorJSON)
			return
		}
		resultJSON, _ := json.Marshal(result)
		res.Write(resultJSON)
	}
}
