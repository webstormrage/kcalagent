package queryProductKPFC

import (
	"ai-kcal-agent/pkg/aiAgent"
	kcaldb "ai-kcal-agent/pkg/kcalDb"
	"encoding/json"
	"fmt"
	"google.golang.org/genai"
)

type JsonProduct struct {
	Name string `json:"name"`
	Kcal int32 `json:"kcal"`
	Proteins int32 `json:"proteins"`
	Fats int32 `json:"fats"`
	Carbohydrates int32 `json:"carbohydrates"`
}

const promptTemplate = "Сколько ккал, белков, жиров, углеводов в 100гр: %s"

func QueryAi(input string) (string, error) {
	prompt := fmt.Sprintf(promptTemplate, input)
	config := &genai.GenerateContentConfig{
        ResponseMIMEType: "application/json",
        ResponseSchema: &genai.Schema{
                Type: genai.TypeObject,
                Properties: map[string]*genai.Schema{
                    "name": {Type: genai.TypeString},
					"kcal": {Type: genai.TypeInteger},
					"proteins": {Type: genai.TypeInteger},
					"fats": {Type: genai.TypeInteger},
					"carbohydrates": {Type: genai.TypeInteger},
				},
                PropertyOrdering: []string{"name", "kcal", "proteins", "fats", "carbohydrates"},
            },
    }

	answer, err :=  aiAgent.Post(prompt, config)
	if err != nil {
		return "", err
	}
	return answer, nil
}

func Parse(data string)(*kcaldb.Product, error) {
	var product JsonProduct
	err := json.Unmarshal([]byte(data), &product)
	if err != nil {
		return nil, err
	}
	return &kcaldb.Product{
		Name: product.Name,
		Kcal: float64(product.Kcal),
		Proteins: float64(product.Proteins),
		Fats: float64(product.Fats),
		Carbohydrates: float64(product.Carbohydrates),
	}, err
}
