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
			Column:     "order_id int UNSIGNED AUTO_INCREMENT, district_id int UNSIGNED NOT NULL, added_on datetime NOT NULL, closed datetime",
			PrimaryKey: "order_id",
			ForeignKey: Foreign{},
		},
		{
			Name: "entries",
			Column: `entry_id int UNSIGNED AUTO_INCREMENT, order_id int UNSIGNED NOT NULL, price decimal(11,2) NOT NULL, 
				status enum('awaiting','inprocess','done','taken') NOT NULL DEFAULT 'awaiting', click_id int UNSIGNED NOT NULL DEFAULT 0`,
			PrimaryKey: "entry_id",
			ForeignKey: Foreign{
				Key:        "order_id",
				References: "orders",
				Column:     "order_id",
			},
		},
		{
			Name: "receipts",
			Column: `receipt_id int UNSIGNED AUTO_INCREMENT, order_id int UNSIGNED NOT NULL, payment enum('card','cash','online') NOT NULL, 
				value decimal(11,2) NOT NULL, accepted_on datetime NOT NULL`,
			PrimaryKey: "receipt_id",
			ForeignKey: Foreign{
				Key:        "order_id",
				References: "orders",
				Column:     "order_id",
			},
		},
	}
}
