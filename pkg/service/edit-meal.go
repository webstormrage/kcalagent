package service

import (
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"encoding/json"
	"net/http"
)

type mealEditReq struct {
	ID            int64   `json:"id"`
	Product       string  `json:"product"`
	Volume        int64   `json:"volume"`
	Kcal          float64 `json:"kcal"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func EditMealHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appContext.Get()
	if r.Method != http.MethodPost {
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

	var req mealEditReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// 3) Обновление meal в базе
	payload := &kcaldb.MealEditPayload{
		ID:            req.ID,
		Name:          req.Product,
		Kcal:          req.Kcal,
		Weight:        float64(req.Volume),
		Proteins:      req.Proteins,
		Fats:          req.Fats,
		Carbohydrates: req.Carbohydrates,
	}
	_, err = kcaldb.UpdateUserMeal(payload, user.ID)
	if err != nil {
		ctx.Logger.Printf("[error]: failed to edit meal\n")
		http.Error(w, "failed to edit meal", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(payload)
	return
}
