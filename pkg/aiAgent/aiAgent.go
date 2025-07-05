package aiAgent

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"ai-kcal-agent/pkg/appContext"
)

func Post(prompt string, config *genai.GenerateContentConfig) (ret string, err error){
	ctx := appContext.Get()
	defer func(){
		if r := recover(); r != nil {
			ret = ""
			err = fmt.Errorf("<panic recovery>: %w", r)
		}
	}()
	context := context.Background()
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey: ctx.GenAiApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	result, err := client.Models.GenerateContent(
        context,
        "gemini-2.5-flash",
        genai.Text(prompt),
        config,
    )
	return result.Text(), err
}
