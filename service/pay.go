package service

import (
	"database/sql"
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

type receiptResult struct {
	res        string
	receipt_id []int64
}

func Pay(db *sql.DB, rowValue ReceiptValue) (receiptResult, error) {
	datetime := time.Now()

	row := db.QueryRow("SELECT SUM(price) FROM entries WHERE ordier_id = ?", rowValue.OrderID)
	var orderPrice float64
	errScan := row.Scan(&orderPrice)

	if errScan != nil {
		return receiptResult{}, errScan
	}

	rowPrice := 0.0

	for _, v := range rowValue.Price {
		rowPrice += v.Value
	}

	if rowPrice != orderPrice {
		return receiptResult{}, errScan
	}

	var receiptID []int64
	for i, v := range rowValue.Price {
		res, errRes := db.Exec("INSERT receipts(order_id, payment, value, accepted_on) VALUES (?,?,?,?)",
			rowValue.OrderID, v.Payment, v.Value, datetime.Format("2006-01-02 15:04:05"))
		if errRes != nil {
			fmt.Printf("failed to insert receips table: %s\n", errRes)
		}
		receiptID[i], _ = res.LastInsertId()
	}

	return receiptResult{res: "success", receipt_id: receiptID}, nil
}
