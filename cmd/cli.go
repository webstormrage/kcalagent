package main

import (
	"ai-kcal-agent/pkg/kcalAi"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"fmt"
	"os"
	"time"
)

func getDayStart() time.Time {
	dayStartTime := time.Now()
	dayStartTime = dayStartTime.Add(-6 * time.Hour)
	dayStartTime = time.Date(dayStartTime.Year(), dayStartTime.Month(), dayStartTime.Day(),
		6, 0, 0, 0, dayStartTime.Location())
	return dayStartTime
}

func aggregate(records []kcalAi.ProductRecord) {
	var kcal float32 = 0
	var proteins float32 = 0
	var fats float32 = 0
	var carbohydrates float32 = 0
	details := ""
	for _, v := range records {
		kcal += v.GetKcal()
		proteins += v.GetProteins()
		fats += v.GetFats()
		carbohydrates += v.GetCarbohydrates()
		details += fmt.Sprintf("%s: %.0f (%.0f, %.0f, %.0f)\n", v.Product.Name, v.GetKcal(), v.GetProteins(), v.GetFats(), v.GetCarbohydrates())
	}
	err := os.WriteFile("logs/multi-output-details", []byte(details), 0644)
	fmt.Printf(details)
	fmt.Printf("Промежуточный итог: %0.f (%0.f, %0.f, %0.f)\n", kcal, proteins, fats, carbohydrates)
	if err != nil {
		fmt.Println(err)
	}
}

func saveToDatabase(records []kcalAi.ProductRecord) {
	for _, v := range records {
		fmt.Printf("[Сохранение в базу]: %s\n", v.Product.Name)
		err := kcaldb.SaveMeals(&kcaldb.MealPayload{
			Name:          v.Product.Name,
			Weight:        float64(v.Weight),
			Kcal:          float64(v.GetKcal()),
			Proteins:      float64(v.GetProteins()),
			Fats:          float64(v.GetFats()),
			Carbohydrates: float64(v.GetCarbohydrates()),
		})
		if err != nil {
			fmt.Println("[Ошибка сохранения в базу]:", err)
		} else {
			fmt.Printf("[Сохранено в базу]: %s\n", v.Product.Name)
		}
	}
}

func printReport() {
	dayStartTime := getDayStart()
	fmt.Println("[Запрос отчета в базу]:", dayStartTime)
	report, err := kcaldb.GetDailySummary(dayStartTime)
	if err != nil {
		fmt.Println("[Ошибка запроса в базу]:", err)
	} else {
		fmt.Printf("Итог за день: %.0f (%.0f, %.0f, %.0f)\n", report.Kcal, report.Proteins, report.Fats, report.Carbohydrates)
	}
}

func runCli() {
	err := kcaldb.SetupDb()
	if err != nil {
		panic(err)
	}
	fileData, err := os.ReadFile("logs/input")
	if err != nil {
		panic(err)
	}
	records, err := kcalAi.GetMealKPFC(string(fileData))
	if err != nil {
		fmt.Println(err)
		return
	}
	aggregate(records)
	var saveToDbAnswer string
	fmt.Println("Сохранить запись в базу?")
	fmt.Scanf("%s\n", &saveToDbAnswer)
	if saveToDbAnswer == "y" {
		saveToDatabase(records)
	}
	printReport()
}
