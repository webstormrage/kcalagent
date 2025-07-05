package kcalAi

import (
	"ai-kcal-agent/pkg/kcalAi/queryMapMeals"
	"ai-kcal-agent/pkg/kcalAi/queryProductKPFC"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"fmt"
	"strings"
)

type ProductRecord struct {
	Product kcaldb.Product
	Weight float32
	LlmResult string;
	Error error;
}

func (r *ProductRecord) GetKcal()float32{
	return r.Weight * float32(r.Product.Kcal) / 100
}

func (r *ProductRecord) GetProteins()float32{
	return r.Weight * float32(r.Product.Proteins) / 100
}

func (r *ProductRecord) GetFats()float32{
	return r.Weight * float32(r.Product.Fats) / 100
}

func (r *ProductRecord) GetCarbohydrates()float32{
	return r.Weight * float32(r.Product.Carbohydrates) / 100
}

func mealItemToProduct(meal *queryMapMeals.MealItem)(*ProductRecord) {
	product, err := kcaldb.GetProductByName(meal.Name)
	fmt.Printf("[Поиск в таблице products]: %s\n", meal.Name)
	if err != nil {
		return &ProductRecord{
			Error: err,
		}
	}
	if product != nil {
		fmt.Printf("[Совпадение в таблице products]: %s\n", product.Name)
		return &ProductRecord{
			Product: *product,
			Weight: float32(meal.Weight),
		}
	}
	fmt.Printf("[Не найдено в таблице products]: %s\n", meal.Name)
	fmt.Printf("[Поиск в таблице products_aliases]: %s\n", meal.Name)
	product, err = kcaldb.GetProductByAlias(meal.Name)
	if err != nil {
		return &ProductRecord{
			Error: err,
		}
	}
	if product != nil {
		fmt.Printf("[Совпадение в таблице products_aliases]: %s\n", product.Name)
		return &ProductRecord{
			Product: *product,
			Weight: float32(meal.Weight),
		}
	}
	fmt.Printf("[Не найдено в таблице products_aliases]: %s\n", meal.Name)
	fmt.Printf("[genai Обработка ввода]: %s\n", meal.Name)
	answer, err := queryProductKPFC.QueryAi(meal.Name)
	fmt.Printf("[genai преобразование в json]:\n%s\n", answer)
	product, err = queryProductKPFC.Parse(answer)
	fmt.Printf("[genai ответ]: в 100гр %s - %.0f ккал\n", product.Name, product.Kcal)
	return &ProductRecord{
		Product: *product,
		Weight: float32(meal.Weight),
		Error: err,
		LlmResult: answer,
	}
}

func mapMealsToProducts(meals []queryMapMeals.MealItem)([]ProductRecord) {
	records := []ProductRecord{}
	for _, meal := range meals {
		recordItem := mealItemToProduct(&meal)
		records = append(records, *recordItem)
	}
	return records
}

func GetMealKPFC(input string)([]ProductRecord, error) {
	formattedInput := strings.ToLower(input)
	fmt.Printf("[genai преобразование в json]:\n%s\n", input)
	response, err := queryMapMeals.QueryAi(formattedInput)
    if err != nil {
		return nil, err
	}
	fmt.Printf("[Парсинг в json]:\n%s\n", response)
	meals, err := queryMapMeals.Parse(response)
	if err != nil {
		return nil, err
	}
	records := mapMealsToProducts(meals)
	return records, nil
}