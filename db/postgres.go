package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/romanova09/bold-trunojoyo-scholarship-api/config"
)

func NewPostgres(cfg *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
