package ghmods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func Chat(token string, model string, messages []Message) (Message, error) {
	url := "https://models.github.io"
	
	payload, err := json.Marshal(map[string]interface{}{
		"model":    model,
		"messages": messages,
	})
	if err != nil {
		return Message{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return Message{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Message{}, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return Message{}, fmt.Errorf("error (%d): %s", resp.StatusCode, string(respData))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respData, &chatResp); err != nil {
		return Message{}, err
	}

	if len(chatResp.Choices) == 0 {
		return Message{}, fmt.Errorf("no choices returned")
	}

	return chatResp.Choices[0].Message, nil
}
