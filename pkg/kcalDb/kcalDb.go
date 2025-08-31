package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func SetupDb() error {
	ctx := appContext.Get()

	var db *sql.DB
	var err error

	// Ретраи подключения: максимум 10 раз с интервалом 3 сек
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", ctx.DataSourceName)
		if err == nil {
			err = db.Ping()
			if err == nil {
				ctx.Logger.Println("✅ Connected to Postgres")
				break
			}
		}
		ctx.Logger.Printf("⏳ Waiting for Postgres... attempt %d/10\n", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("could not connect to Postgres: %w", err)
	}
	defer db.Close()

	// Применяем schema-1.sql
	fileData, err := os.ReadFile("sql/schema-1.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = db.Exec(string(fileData))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}
