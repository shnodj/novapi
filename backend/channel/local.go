package channel

import (
	"context"
	"fmt"
	"novapi/backend/model"
)

// RunLocal handles local binary/Ollama requests
func RunLocal(ctx context.Context, channel *model.Channel, prompt string) (string, error) {
	// Implementation for running local binary or calling local Ollama
	fmt.Printf("Running local model: %s\n", channel.Name)
	return "Mock response from local model", nil
}
