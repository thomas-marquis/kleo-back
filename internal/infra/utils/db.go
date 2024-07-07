package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thomas-marquis/kleo-back/internal/infra/config"
)

func NewDB(dbConfig config.DatabaseConfig) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database))
	if err != nil {
		errorMsg := fmt.Sprintf("Error while connecting to database: %s", err.Error())
		panic(errorMsg)
	}
	return db
}
