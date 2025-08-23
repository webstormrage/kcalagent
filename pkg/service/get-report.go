package service

import (
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"encoding/json"
	"net/http"
	"time"
)

func getDayStart(offsetHours int) time.Time {
	loc := time.FixedZone("CustomTZ", offsetHours*3600)

	// берём текущее время в этой зоне
	dayStartTime := time.Now().In(loc)
	dayStartTime = dayStartTime.Add(-6 * time.Hour)
	dayStartTime = time.Date(dayStartTime.Year(), dayStartTime.Month(), dayStartTime.Day(),
		6, 0, 0, 0, loc)
	return dayStartTime
}

func GetDailyReportHandler(w http.ResponseWriter, r *http.Request) {
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

	// 3) Начало дня (пример с UTC+4)
	dayStartTime := getDayStart(4)

	// 4) Получаем суточную сводку для пользователя
	report, err := kcaldb.GetDailySummaryByUser(dayStartTime, user.ID)
	if err != nil {
		http.Error(w, "failed to get report", http.StatusInternalServerError)
		return
	}

	// 5) Отдаём JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(report); err != nil {
		http.Error(w, "failed to get report", http.StatusInternalServerError)
	}
}
