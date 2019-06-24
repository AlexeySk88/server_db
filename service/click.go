package service

import (
	"database/sql"
)

type ClickValue struct {
	EntryID int
}

type clickResult struct {
	entryID int
	status  string
}

func Click(db *sql.DB, rowValue ClickValue) (clickResult, error) {
	_, errRes := db.Exec(`
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
		return clickResult{}, errRes
	}

	res := clickResult{}
	row := db.QueryRow("SELECT entry_id, status FROM entries WHERE entry_id = ?", rowValue.EntryID)
	errScan := row.Scan(&res.entryID, &res.status)

	if errScan != nil {
		return clickResult{}, errRes
	}

	return res, nil
}
