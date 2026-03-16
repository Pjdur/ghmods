# ghmods: A Go client for GitHub-hosted AI models

**ghmods** is a Go client for accessing GitHub-hosted AI models via the `https://models.github.ai` API. It provides a simple interface for chat-based completions, similar to OpenAI’s API, but powered by GitHub’s model infrastructure.

---

## ✨ Features

- Chat completion support for GitHub-hosted models (e.g., `openai/gpt-4o`)
- Easy authentication via GitHub personal access token (PAT)
- Clean and ergonomic API

---

## 🚀 Getting Started

### 1. Install

Add it to your Go project:

```bash
go get github.com/Pjdur/ghmods
```

### 2. Authenticate

Set your GitHub token as an environment variable:

```bash
export GITHUB_TOKEN=your_personal_access_token
```

> 🔐 Your token must have the **Models** permission enabled.  
> You can generate a PAT in your GitHub settings.
> [Managing your personal access tokens](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)

---

## Basic Example (Single-Turn)

Here's a simple example demonstrating how to use `ghmods` for a single chat completion:

```go /dev/null/example.go#L1-28
package main

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
```

To run this example, save it as `main.go` in your project and execute:

```bash
go run main.go
```

---

## 📚 API Overview

### `ghmods.Chat(token string, model string, messages []ghmods.Message) (ghmods.Message, error)`

Sends a chat request to the model endpoint.

**Parameters:**
- `token`: Your GitHub personal access token.
- `model`: The model name (e.g., `"openai/gpt-4o"`).
- `messages`: A slice of `ghmods.Message` structs representing the conversation history.

**Returns:**
- `ghmods.Message`: The assistant's reply.
- `error`: An error if the request fails or the response is invalid.

---

## 🛠️ Development

To contribute or run examples locally:

```bash
git clone https://github.com/Pjdur/ghmods
cd ghmods
```

---

## 📄 License

MIT © Pjdur
