package routed

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server_db/service"
)

type ErrorRes struct {
	Res    string
	Report string
}

func handleAddOrder(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		body, errRead := ioutil.ReadAll(req.Body)
		if errRead != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errRead.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		reqJSON := []byte(body)
		order := service.OrderValue{}
		errType := json.Unmarshal(reqJSON, &order)
		if errType != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errType.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}

		result, err := service.Order(db, order)
		if err != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: err.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		resultJSON, errMarshal := json.Marshal(result)
		if errMarshal != nil {
			log.Println("marchal error: ", errMarshal)
		}
		res.Write(resultJSON)
	}
}

func handleAddReceipt(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		body, errRead := ioutil.ReadAll(req.Body)
		if errRead != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errRead.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		reqJSON := []byte(body)
		receipt := service.ReceiptValue{}
		errType := json.Unmarshal(reqJSON, &receipt)
		if errType != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errType.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}

		result, err := service.Pay(db, receipt)
		if err != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: err.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		resultJSON, errMarshal := json.Marshal(result)
		if errMarshal != nil {
			log.Println("marchal error: ", errMarshal)
		}
		res.Write(resultJSON)
	}
}

func handlerClick(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		body, errRead := ioutil.ReadAll(req.Body)
		if errRead != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errRead.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		reqJSON := []byte(body)
		click := service.ClickValue{}
		errType := json.Unmarshal(reqJSON, &click)
		if errType != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errType.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}

		result, err := service.Click(db, click)
		if err != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: err.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		resultJSON, errMarshal := json.Marshal(result)
		if errMarshal != nil {
			log.Println("marchal error: ", errMarshal)
		}
		res.Write(resultJSON)
	}
}

func handlerDelivered(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		body, errRead := ioutil.ReadAll(req.Body)
		if errRead != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errRead.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		reqJSON := []byte(body)
		deliv := service.DeliveryValue{}
		errType := json.Unmarshal(reqJSON, &deliv)
		if errType != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: errType.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}

		result, err := service.Delivered(db, deliv)
		if err != nil {
			errorJSON, errMarshal := json.Marshal(ErrorRes{Res: "fail", Report: err.Error()})
			if errMarshal != nil {
				log.Println("marchal error: ", errMarshal)
			}
			res.Write(errorJSON)
			return
		}
		resultJSON, errMarshal := json.Marshal(result)
		if errMarshal != nil {
			log.Println("marchal error: ", errMarshal)
		}
		res.Write(resultJSON)
	}
}
