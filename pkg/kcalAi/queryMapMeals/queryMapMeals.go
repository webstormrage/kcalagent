package queryMapMeals

import (
	"ai-kcal-agent/pkg/aiAgent"
	"google.golang.org/genai"
	"encoding/json"
	"fmt"
)

type MealItem struct {
	Name string `json:"name"`
	Weight int32 `json:"weight"`
}

const promptTemplate = "Преобразуй эти данные о названии продукта и весе в json согласно схеме: \n%s"

func QueryAi(input string)(string, error) {
	config := &genai.GenerateContentConfig{
        ResponseMIMEType: "application/json",
        ResponseSchema: &genai.Schema{
            Type: genai.TypeArray,
            Items: &genai.Schema{
                Type: genai.TypeObject,
                Properties: map[string]*genai.Schema{
                    "name": {Type: genai.TypeString},
					"weight": {Type: genai.TypeInteger},
				},
                PropertyOrdering: []string{"name", "weight"},
            },
        },
    }
   return aiAgent.Post(fmt.Sprintf(promptTemplate, input), config)
}

func Parse(data string)([]MealItem, error) {
	var items []MealItem
	err := json.Unmarshal([]byte(data), &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}