package service

import (
	"ai-kcal-agent/pkg/aiAgent"
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"encoding/json"
	"net/http"
)

type mealReq struct {
	Product    string `json:"product"`
	Volume     int64  `json:"volume"`
	GenAiToken string `json:"genAIToken"`
}

type mealResp struct {
	ID            int64   `json:"id"`
	Product       string  `json:"product"`
	Volume        int64   `json:"volume"`
	Kcal          float64 `json:"kcal"`
	Proteins      float64 `json:"proteins"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
	Strategy      string  `json:"strategy"`
}

func productToMeal(product *kcaldb.Product, volume int64, strategy string) *mealResp {
	return &mealResp{
		Product:       product.Name,
		Volume:        volume,
		Kcal:          product.Kcal * float64(volume) / 100,
		Proteins:      product.Proteins * float64(volume) / 100,
		Fats:          product.Fats * float64(volume) / 100,
		Carbohydrates: product.Carbohydrates * float64(volume) / 100,
		Strategy:      strategy,
	}
}

func AddMealHandler(w http.ResponseWriter, r *http.Request) {
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

	var req mealReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	var resp *mealResp

	// 3) Поиск продукта
	product, err := kcaldb.GetUserProductByName(req.Product, user.ID)
	if err != nil {
		ctx.Logger.Printf("[error]: get product failed: %v\n", err)
	}
	if product != nil {
		resp = productToMeal(product, req.Volume, "exact")
	}

	// 4) Поиск продукта по алиасу
	if product == nil {
		product, err = kcaldb.GetUserProductByAlias(req.Product, user.ID)
		if err != nil {
			ctx.Logger.Printf("[error]: get alias failed: %v\n", err)
		}
		if product != nil {
			resp = productToMeal(product, req.Volume, "alias")
		}
	}

	// 4) Поиск продукта в genai
	if product == nil {
		modelResp, err := aiAgent.QueryProduct(req.Product, req.GenAiToken)
		if err != nil {
			ctx.Logger.Printf("[error]: query ai product failed: %v\n", err)
			http.Error(w, "ai request failed", http.StatusInternalServerError)
			return
		}
		// 5) Парсинг genai ответа
		product, err = aiAgent.ParseProduct(modelResp)
		if err != nil {
			ctx.Logger.Printf("[error]: query ai reponse parsing failed: %v\n", err)
			http.Error(w, "ai response parsing failed", http.StatusInternalServerError)
			return
		}
		if product != nil {
			resp = productToMeal(product, req.Volume, "genai")
		}
	}

	if resp == nil {
		ctx.Logger.Printf("[error]: failed to get meal\n")
		http.Error(w, "failed to get meal", http.StatusInternalServerError)
	}
	// 6) Сохранение meal в базу
	mealId, err := kcaldb.SaveUserMeal(&kcaldb.MealPayload{
		Name:          resp.Product,
		Kcal:          resp.Kcal,
		Weight:        float64(resp.Volume),
		Proteins:      resp.Proteins,
		Fats:          resp.Fats,
		Carbohydrates: resp.Carbohydrates,
	}, user.ID)
	if err != nil {
		ctx.Logger.Printf("[error]: failed to save meal\n")
		http.Error(w, "failed to save meal", http.StatusInternalServerError)
	}
	resp.ID = mealId

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
	return
}
