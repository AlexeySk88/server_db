package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type DeliveryValue struct {
	OrderID int
}

type DeliveryResult struct {
	Res      string
	Order_id int
}

func Delivered(db *sql.DB, rowValue DeliveryValue) (DeliveryResult, error) {
	res, errRes := db.Exec(
		`UPDATE entries AS e 
		JOIN orders AS o ON o.order_id = e.order_id 
		SET status='taken', closed=NOW()
		WHERE e.order_id = ?`,
		rowValue.OrderID)
	if errRes != nil {
		return DeliveryResult{}, errRes
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Println("click error: ", err)
	}
	if count == 0 {
		return DeliveryResult{}, errors.New(fmt.Sprint("not found order with order_id=", rowValue.OrderID))
	}

	return DeliveryResult{Res: "success", Order_id: rowValue.OrderID}, nil
}
