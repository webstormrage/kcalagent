package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
)

type MealPayload struct {
	Name          string
	Weight        float64
	Kcal          float64
	Proteins      float64
	Fats          float64
	Carbohydrates float64
}

func SaveUserMeal(meal *MealPayload, userId int64) (int64, error) {
	app := appContext.Get()

	db, err := sql.Open("postgres", app.DataSourceName)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	app.Logger.Printf("[DB]: save meal query <%s> <%d>\n", meal.Name, userId)

	var id int64
	err = db.QueryRow(
		`INSERT INTO meals (name, weight, kcal, proteins, fats, carbohydrates, user_id)
		 VALUES ($1,   $2,     $3,   $4,       $5,   $6,             $7)
		 RETURNING id`,
		meal.Name,
		meal.Weight,
		meal.Kcal,
		meal.Proteins,
		meal.Fats,
		meal.Carbohydrates,
		userId,
	).Scan(&id)
	if err != nil {
		app.Logger.Printf("[DB]: save meal error <%s> <%d>: %v\n", meal.Name, userId, err)
		return 0, err
	}

	app.Logger.Printf("[DB]: save meal success id=%d <%s> <%d>\n", id, meal.Name, userId)
	return id, nil
}

type MealEditPayload struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Weight        float64 `json:"weight"`
	Kcal          float64 `json:"kcal"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func UpdateUserMeal(meal *MealEditPayload, userId int64) (int64, error) {
	app := appContext.Get()

	db, err := sql.Open("postgres", app.DataSourceName)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	app.Logger.Printf("[DB]: update meal id=%d user_id=%d <%s>\n", meal.ID, userId, meal.Name)

	var id int64
	err = db.QueryRow(
		`UPDATE meals
		    SET name          = $1,
		        weight        = $2,
		        kcal          = $3,
		        proteins      = $4,
		        fats          = $5,
		        carbohydrates = $6
		  WHERE id = $7 AND user_id = $8
		  RETURNING id`,
		meal.Name,
		meal.Weight,
		meal.Kcal,
		meal.Proteins,
		meal.Fats,
		meal.Carbohydrates,
		meal.ID,
		userId,
	).Scan(&id)

	if err == sql.ErrNoRows {
		app.Logger.Printf("[DB]: update meal not found id=%d user_id=%d\n", meal.ID, userId)
		return 0, err
	}
	if err != nil {
		app.Logger.Printf("[DB]: update meal error id=%d user_id=%d: %v\n", meal.ID, userId, err)
		return 0, err
	}

	app.Logger.Printf("[DB]: update meal success id=%d user_id=%d\n", id, userId)
	return id, nil
}

func GetUserMeal(mealID, userId int64) (*MealEditPayload, error) {
	app := appContext.Get()

	db, err := sql.Open("postgres", app.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	app.Logger.Printf("[DB]: get meal id=%d user_id=%d\n", mealID, userId)

	m := &MealEditPayload{}
	err = db.QueryRow(
		`SELECT id, name, weight, kcal, proteins, fats, carbohydrates
		   FROM meals
		  WHERE id = $1 AND user_id = $2`,
		mealID,
		userId,
	).Scan(
		&m.ID,
		&m.Name,
		&m.Weight,
		&m.Kcal,
		&m.Proteins,
		&m.Fats,
		&m.Carbohydrates,
	)

	if err == sql.ErrNoRows {
		app.Logger.Printf("[DB]: meal not found id=%d user_id=%d\n", mealID, userId)
		return nil, err
	}
	if err != nil {
		app.Logger.Printf("[DB]: get meal error id=%d user_id=%d: %v\n", mealID, userId, err)
		return nil, err
	}

	app.Logger.Printf("[DB]: get meal success id=%d user_id=%d <%s>\n", m.ID, userId, m.Name)
	return m, nil
}
