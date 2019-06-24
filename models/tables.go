package models

type DistrictTable struct {
	Name       string
	Column     string
	PrimaryKey string
	ForeignKey Foreign
}

type Foreign struct {
	Key        string
	References string
	Column     string
}

func GetDistrictTable() []DistrictTable {
	return []DistrictTable{
		{
			Name:       "orders",
			Column:     "order_id int UNSIGNED AUTO_INCREMENT, district_id int UNSIGNED, added_on datetime, closed datetime",
			PrimaryKey: "order_id",
			ForeignKey: Foreign{},
		},
		{
			Name:       "entries",
			Column:     "entry_id int UNSIGNED AUTO_INCREMENT, order_id int UNSIGNED, price decimal(11,2), status enum('awaiting','inprocess','done','taken') DEFAULT 'awaiting', click_id int UNSIGNED DEFAULT 0",
			PrimaryKey: "entry_id",
			ForeignKey: Foreign{
				Key:        "order_id",
				References: "orders",
				Column:     "order_id",
			},
		},
		{
			Name:       "receipts",
			Column:     "receipt_id int UNSIGNED AUTO_INCREMENT, order_id int UNSIGNED, payment enum('card','cash','online'), value decimal(11,2), accepted_on datetime",
			PrimaryKey: "receipt_id",
			ForeignKey: Foreign{
				Key:        "order_id",
				References: "orders",
				Column:     "order_id",
			},
		},
	}
}
