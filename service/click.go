package service

import (
	"database/sql"
	"errors"
	"log"
)

type ClickValue struct {
	EntryID int
}

type ClickResult struct {
	EntryID int
	Status  string
}

func Click(db *sql.DB, rowValue ClickValue) (ClickResult, error) {
	upRes, errRes := db.Exec(`
		UPDATE entries AS e1
		JOIN (
			SELECT click_id
			FROM entries
			WHERE entry_id = ?
			) AS e2
		SET e1.status = IF(e2.click_id = 0, 'inprocess', 'done'), e1.click_id = e1.click_id+1
		WHERE e1.entry_id = ? AND e1.click_id < 2`,
		rowValue.EntryID, rowValue.EntryID)
	if errRes != nil {
		return ClickResult{}, errRes
	}

	count, err := upRes.RowsAffected()
	if err != nil {
		log.Println("click error: ", err)
	}
	if count == 0 {
		return ClickResult{}, errors.New("Update data not found")
	}

	res := ClickResult{}
	row := db.QueryRow("SELECT entry_id, status FROM entries WHERE entry_id = ?", rowValue.EntryID)
	errScan := row.Scan(&res.EntryID, &res.Status)

	if errScan != nil {
		return ClickResult{}, errRes
	}

	return res, nil
}
