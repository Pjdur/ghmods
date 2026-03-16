package ghmods

import (
	"fmt"
	"os"

	"github.com/Pjdur/ghmods"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Error: GITHUB_TOKEN environment variable not set.")
		return
	}

	messages := []ghmods.Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: "What is the capital of France?"},
	}

	response, err := ghmods.Chat(token, "openai/gpt-4o", messages)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(response.Content)
}