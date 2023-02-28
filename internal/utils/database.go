package utils

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.name")))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %v", err)
	}

	return db, nil
}
