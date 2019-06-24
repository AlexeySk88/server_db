package routed

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"server_db/service"
)

type errorRes struct {
	res    string
	report error
}

func handleAddOrder(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		reqJSON := []byte(req.URL.Query().Get("order"))
		order := service.OrderValue{}
		errType := json.Unmarshal(reqJSON, &order)
		if errType != nil {
			errorJSON, _ := json.Marshal(errorRes{res: "fail", report: errType})
			fmt.Fprintf(res, string(errorJSON))
		}

		_, err := service.Order(db, order)
		if err != nil {
			//errorJSON, _ := json.Marshal(errorRes{res: "fail", report: err})
			fmt.Fprintf(res, "fail")
		}

		fmt.Fprintf(res, "OK1")
	}
}

func handleAddReceipt(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		reqJSON := []byte(req.URL.Query().Get("receipt"))
		receipt := service.ReceiptValue{}
		errType := json.Unmarshal(reqJSON, &receipt)
		if errType != nil {
			errorJSON, _ := json.Marshal(errorRes{res: "fail", report: errType})
			fmt.Fprintf(res, string(errorJSON))
		}

		result, err := service.Pay(db, receipt)
		if err != nil {
			errorJSON, _ := json.Marshal(errorRes{res: "fail", report: err})
			fmt.Fprintf(res, string(errorJSON))
		}

		resJSON, _ := json.Marshal(result)
		fmt.Fprintf(res, string(resJSON))
	}
}

func handlerClick(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		reqJSON := []byte(req.URL.Query().Get("click"))
		click := service.ClickValue{}
		errType := json.Unmarshal(reqJSON, &click)
		if errType != nil {
			errorJSON, _ := json.Marshal(errorRes{res: "fail", report: errType})
			fmt.Fprintf(res, string(errorJSON))
		}

		result, err := service.Click(db, click)
		if err != nil {
			errorJSON, _ := json.Marshal(errorRes{res: "fail", report: err})
			fmt.Fprintf(res, string(errorJSON))
		}

		resJSON, _ := json.Marshal(result)
		fmt.Fprintf(res, string(resJSON))
	}
}
