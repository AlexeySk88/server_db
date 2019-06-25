package service

import (
	"database/sql"
)

type DeliveryValue struct {
	OrderID int
}

type DeliveryResult struct {
	Res      string
	Order_id int
}

func Delivered(db *sql.DB, rowValue DeliveryValue) (DeliveryResult, error) {
	_, errRes := db.Exec(
		`UPDATE entries AS e 
		JOIN orders AS o ON o.order_id = e.order_id 
		SET status='taken', closed=NOW()
		WHERE e.order_id = ?`,
		rowValue.OrderID)
	if errRes != nil {
		return DeliveryResult{}, errRes
	}

	return DeliveryResult{Res: "success", Order_id: rowValue.OrderID}, nil
}
