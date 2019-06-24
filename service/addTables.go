package service

import (
	"database/sql"
	"fmt"
	"server_db/models"
)

func AddTables(db *sql.DB) {
	districvalues := models.GetDistrictTable()

	for _, value := range districvalues {
		var foreignKey string

		if (value.ForeignKey != models.Foreign{}) {
			foreignKey = fmt.Sprintf(`, FOREIGN KEY (%s)  REFERENCES %s (%s)`, value.ForeignKey.Key, value.ForeignKey.References, value.ForeignKey.Column)
		}

		_, err := db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		%s,
		PRIMARY KEY (%s)
		%s
		)`,
			value.Name, value.Column, value.PrimaryKey, foreignKey))

		if err != nil {
			fmt.Printf("failed to create table %s, %s\n", value.Name, err)
		}
	}
}
