package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type ReceiptValue struct {
	OrderID int
	Price   []struct {
		Payment string
		Value   float64
	}
}

type ReceiptResult struct {
	Res        string
	Receipt_id []int64
}

func Pay(db *sql.DB, rowValue ReceiptValue) (ReceiptResult, error) {
	datetime := time.Now()

	row := db.QueryRow("SELECT SUM(price) FROM entries WHERE order_id = ?", rowValue.OrderID)
	var orderPrice float64
	errScan := row.Scan(&orderPrice)

	if errScan != nil {
		return ReceiptResult{}, errors.New(fmt.Sprint("not found order with order_id=", rowValue.OrderID))
	}

	rowPrice := 0.0

	for _, v := range rowValue.Price {
		rowPrice += v.Value
	}

	if rowPrice != orderPrice {
		return ReceiptResult{}, errors.New("Enter the total price of the order")
	}

	var receiptID []int64
	for _, v := range rowValue.Price {
		res, errRes := db.Exec("INSERT receipts(order_id, payment, value, accepted_on) VALUES (?,?,?,?)",
			rowValue.OrderID, v.Payment, v.Value, datetime.Format("2006-01-02 15:04:05"))
		if errRes != nil {
			return ReceiptResult{}, errRes
		}
		id, _ := res.LastInsertId()
		receiptID = append(receiptID, id)
	}

	return ReceiptResult{Res: "success", Receipt_id: receiptID}, nil
}
