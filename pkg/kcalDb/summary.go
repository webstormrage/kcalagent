package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
	"time"
)

type MealSummary struct {
	Kcal          float64 `json:"kcal"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func GetDailySummaryByUser(dayStartTime time.Time, userId int64) (*MealSummary, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ctx.Logger.Printf("[DB]: get daily summary query <%s>\n", userId)
	row := db.QueryRow(
		"SELECT SUM(kcal) AS kcal, SUM(proteins) AS proteins, SUM(fats) AS fats, SUM(carbohydrates) AS carbohydrates FROM meals"+
			" WHERE created_at > $1 AND user_id = $2",
		dayStartTime,
		userId,
	)
	var summary MealSummary
	err = row.Scan(&summary.Kcal, &summary.Proteins, &summary.Fats, &summary.Carbohydrates)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	ctx.Logger.Printf("[DB]: get user by login query success <%s>\n", userId)
	return &summary, err
}
