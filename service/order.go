package service

import (
	"database/sql"
	"time"
)

type OrderValue struct {
	DistrictId int
	Price      []float64
}

type OrderResult struct {
	Res      string
	Order_id int64
	Entry_id []int64
}

func Order(db *sql.DB, rowValue OrderValue) (OrderResult, error) {
	datetime := time.Now()
	resOrder, errOrder := db.Exec("INSERT orders(district_id, added_on) VALUES (?, ?)",
		rowValue.DistrictId, datetime.Format("2006-01-02 15:04:05"))

	if errOrder != nil {
		return OrderResult{}, errOrder
	}

	orderID, _ := resOrder.LastInsertId()

	var entryID []int64
	for _, v := range rowValue.Price {
		entryRes, errEntr := db.Exec("INSERT entries(order_id, price) VALUES (?,?)",
			orderID, v)

		if errEntr != nil {
			return OrderResult{}, errEntr
		}

		id, _ := entryRes.LastInsertId()
		entryID = append(entryID, id)
	}

	return OrderResult{Res: "success", Order_id: orderID, Entry_id: entryID}, nil
}
