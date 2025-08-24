package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
)

func SaveUserMeal(meal *MealPayload, userId int64) error {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	ctx.Logger.Printf("[DB]: save meal query <%s> <%d>\n", meal.Name, userId)
	_, err = db.Exec(
		"INSERT INTO meals (name, weight, kcal, proteins, fats, carbohydrates, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		meal.Name,
		meal.Weight,
		meal.Kcal,
		meal.Proteins,
		meal.Fats,
		meal.Carbohydrates,
		userId,
	)
	if err == nil {
		ctx.Logger.Printf("[DB]: save meal query success <%s> <%d>\n", meal.Name, userId)
	}
	return err
}
