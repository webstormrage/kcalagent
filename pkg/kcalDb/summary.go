package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
	"time"
)

type MealSummary struct {
	Name          string  `json:"name"`
	Kcal          float64 `json:"kcal"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func GetDailySummaryByUser(dayStartTime time.Time, userId int64) ([]MealSummary, error) {
	ctx := appContext.Get()

	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ctx.Logger.Printf("[DB] daily summary user_id=%d since=%s", userId, dayStartTime.Format(time.RFC3339))

	// Все строки + итог, вычисленный в SQL
	const q = `
SELECT name, kcal, proteins, fats, carbohydrates
FROM (
    -- отдельные приёмы
    SELECT
        name,
        COALESCE(kcal, 0)          AS kcal,
        COALESCE(proteins, 0)      AS proteins,
        COALESCE(fats, 0)          AS fats,
        COALESCE(carbohydrates, 0) AS carbohydrates,
        created_at,
        0 AS ord
    FROM meals
    WHERE created_at > $1 AND user_id = $2

    UNION ALL

    -- итоговая строка
    SELECT
        'ИТОГ' AS name,
        COALESCE(SUM(kcal), 0)          AS kcal,
        COALESCE(SUM(proteins), 0)      AS proteins,
        COALESCE(SUM(fats), 0)          AS fats,
        COALESCE(SUM(carbohydrates), 0) AS carbohydrates,
        NULL::timestamp                  AS created_at,
        1 AS ord
    FROM meals
    WHERE created_at > $1 AND user_id = $2
) t
ORDER BY ord DESC, created_at DESC;
`

	rows, err := db.Query(q, dayStartTime, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []MealSummary
	for rows.Next() {
		var m MealSummary
		if err := rows.Scan(&m.Name, &m.Kcal, &m.Proteins, &m.Fats, &m.Carbohydrates); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	ctx.Logger.Printf("[DB] daily summary ok: rows=%d user_id=%d", len(list), userId)
	return list, nil
}
