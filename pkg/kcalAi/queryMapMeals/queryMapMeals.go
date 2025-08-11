package queryMapMeals

import (
	"ai-kcal-agent/pkg/aiAgent"
	"encoding/json"
	"fmt"
	"google.golang.org/genai"
	"strings"
)

type MealValues struct {
	Kcal          int32 `json:"kcal"`
	Proteins      int32 `json:"proteins"`
	Fats          int32 `json:"fats"`
	Carbohydrates int32 `json:"carbohydrates"`
}

type MealItem struct {
	Name    string      `json:"name"`
	Weight  int32       `json:"weight"`
	AddFlag bool        `json:"addFlag"`
	Alias   string      `json:"alias"`
	Values  *MealValues `json:"values,omitempty"`
}

const promptTemplate = "Преобразуй эти данные о названии продукта и весе в json согласно схеме," +
	"если в конце строки есть +, то addFlag=true иначе false" +
	"если в строке есть 4 числа в скобочках, то их нужно соответственно сохранить в поле values.Kcal, values.Proteins, values.Fats, values.Carbohydrates" +
	"иначе в values не должен приходить" +
	"если в скобочках указана строка, то она должна быть сохранено в name, а строка перед скобочками в alias" +
	"если скобочек нет, или в них указаны числа, то alias должен быть пустой строкой" +
	": \n%s"

func QueryAi(input string) (string, error) {
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"name":    {Type: genai.TypeString},
					"weight":  {Type: genai.TypeInteger},
					"addFlag": {Type: genai.TypeBoolean},
					"alias":   {Type: genai.TypeString},
					"values": {
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"kcal":          {Type: genai.TypeInteger},
							"proteins":      {Type: genai.TypeInteger},
							"fats":          {Type: genai.TypeInteger},
							"carbohydrates": {Type: genai.TypeInteger},
						},
					},
				},
				PropertyOrdering: []string{"name", "alias", "weight", "addFlag", "values"},
			},
		},
	}
	return aiAgent.Post(fmt.Sprintf(promptTemplate, input), config)
}

func Parse(data string) ([]MealItem, error) {
	var items []MealItem
	err := json.Unmarshal([]byte(data), &items)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		item.Name = strings.TrimSpace(strings.ToLower(item.Name))
	}
	return items, nil
}
