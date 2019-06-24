package service

import (
	"database/sql"
)

type DeliveryValue struct {
	OrderID int
}

type deliveryResult struct {
	res      string
	order_id int
}

func Delivered(db *sql.DB, rowValue DeliveryValue) (deliveryResult, error) {
	_, errRes := db.Exec("UPDATE entries SET(status = 'taken' WHERE order_id = ?",
		rowValue.OrderID)
	if errRes != nil {
		return deliveryResult{}, errRes
	}

	return deliveryResult{res: "success", order_id: rowValue.OrderID}, nil
}
