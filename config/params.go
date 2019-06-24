package config

import "os"

func Set() {
	os.Setenv("DRIVER", "mysql")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_DATABASE", "compose")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	return
}
