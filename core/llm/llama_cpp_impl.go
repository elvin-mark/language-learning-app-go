package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type llamaCppImpl struct {
	baseUrl string
	model   string
}

func NewLlamaCpp(baseUrl string, model string) Llm {
	return &llamaCppImpl{
		baseUrl: baseUrl,
		model:   model,
	}
}

func (llm *llamaCppImpl) GetResponse(prompt string) (result ChatResponse, err error) {
	// Prepare request
	reqBody := ChatRequest{
		Model:  llm.model,
		Stream: false,
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	// Send POST request
	resp, err := http.Post(
		llm.baseUrl+"/v1/chat/completions",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Parse JSON
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
