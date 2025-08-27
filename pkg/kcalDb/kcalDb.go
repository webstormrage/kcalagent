package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func SetupDb() error {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	fileData, err := os.ReadFile("sql/schema-1.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(fileData))
	return err
}
