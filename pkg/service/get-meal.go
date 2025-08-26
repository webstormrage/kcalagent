package service

import (
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type mealGetResp struct {
	ID            int64   `json:"id"`
	Product       string  `json:"product"`
	Volume        int64   `json:"volume"`
	Kcal          float64 `json:"kcal"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func GetMealHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appContext.Get()
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1) JWT из заголовка X-Auth
	token := r.Header.Get("X-Auth")
	login, err := ParseJWT(token)
	if err != nil {
		ctx.Logger.Printf("[error]: parse jwt failed: %v\n", err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// 2) Находим пользователя
	user, err := kcaldb.GetUserByLogin(login)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	// 3) Читаем id из query (?id=123)
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	mealID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// 4) Получаем meal из БД (учитывается user_id)
	meal, err := kcaldb.GetUserMeal(mealID, user.ID)
	if err == sql.ErrNoRows {
		http.Error(w, "meal not found", http.StatusNotFound)
		return
	}
	if err != nil {
		ctx.Logger.Printf("[error]: get meal failed: %v\n", err)
		http.Error(w, "failed to get meal", http.StatusInternalServerError)
		return
	}

	resp := mealGetResp{
		ID:            meal.ID,
		Product:       meal.Name,
		Volume:        int64(meal.Weight), // Weight -> Volume (округляем вниз до int64)
		Kcal:          meal.Kcal,
		Proteins:      meal.Proteins,
		Fats:          meal.Fats,
		Carbohydrates: meal.Carbohydrates,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
