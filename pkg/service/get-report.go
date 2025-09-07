package service

import (
	"ai-kcal-agent/pkg/appContext"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"encoding/json"
	"net/http"
	"time"
)

func dayWindowUTC(dayOffset int, hoursOffset int) (time.Time, time.Time) {
	if dayOffset < 0 {
		dayOffset = 0
	}
	const dayStartHourLocal = 6 // фиксированное начало суток в локальном времени

	loc := time.FixedZone("CustomTZ", hoursOffset*3600)

	nowLocal := time.Now().In(loc)

	// Если текущее локальное время до 6 утра → считаем, что день ещё "вчерашний"
	if nowLocal.Hour() < dayStartHourLocal {
		dayOffset++
	}

	// Берём локальную дату с учётом сдвига dayOffset
	targetLocal := nowLocal.AddDate(0, 0, -dayOffset)

	// Начало дня: 06:00 локального времени
	startLocal := time.Date(
		targetLocal.Year(), targetLocal.Month(), targetLocal.Day(),
		dayStartHourLocal, 0, 0, 0, loc,
	)

	startUTC := startLocal.UTC()
	endUTC := startUTC.Add(24 * time.Hour)
	return startUTC, endUTC
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
	startUtc, endUtc := dayWindowUTC(0, 4)

	// 4) Получаем суточную сводку для пользователя
	report, err := kcaldb.GetDailySummaryByUser(startUtc, endUtc, user.ID)
	if err != nil {
		ctx.Logger.Printf("[error]: failed to get report: %v\n", err)
		http.Error(w, "failed to get report", http.StatusInternalServerError)
		return
	}

	// 5) Отдаём JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(report); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
