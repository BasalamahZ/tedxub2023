package config

import "os"

func BaseConfig() string {
	return "" +
		"host=" + os.Getenv("PGHOST") +
		" user=" + os.Getenv("PGUSER") +
		" password=" + os.Getenv("PGPASSWORD") +
		" dbname=" + os.Getenv("PGDATABASE") +
		" port=" + os.Getenv("PGPORT") +
		" sslmode=disable"
}
