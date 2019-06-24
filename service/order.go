package service

import (
	"database/sql"
	"time"
)

type OrderValue struct {
	DistrictId int
	Price      float32
}

type orderResult struct {
	res      string
	order_id int64
	entry_id int64
}

func Order(db *sql.DB, rowValue OrderValue) (orderResult, error) {
	datetime := time.Now()
	orderRes, errRes := db.Exec("INSERT orders(district_id, added_on) VALUES (?, ?)",
		rowValue.DistrictId, datetime.Format("2006-01-02 15:04:05"))

	if errRes != nil {
		return orderResult{}, errRes
	}

	orderID, _ := orderRes.LastInsertId()

	entryRes, errEntr := db.Exec("INSERT enties(order_id, price, status) VALUES (?, ?)",
		orderID, rowValue.Price)

	if errEntr != nil {
		return orderResult{}, errRes
	}

	entryID, _ := entryRes.LastInsertId()

	return orderResult{res: "success", order_id: orderID, entry_id: entryID}, nil
}
