package aiAgent

import (
	"ai-kcal-agent/pkg/appContext"
	"context"
	"fmt"
	"google.golang.org/genai"
)

func QueryAi(prompt string, config *genai.GenerateContentConfig, token string) (ret string, err error) {
	defer func() {
		if r := recover(); r != nil {
			ret = ""
			err = fmt.Errorf("<panic recovery>: %w", r)
		}
	}()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  token,
		Backend: genai.BackendGeminiAPI,
	})
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		config,
	)
	return result.Text(), err
}

func Post(prompt string, config *genai.GenerateContentConfig) (ret string, err error) {
	ctx := appContext.Get()
	return QueryAi(prompt, config, ctx.GenAiApiKey)
}
