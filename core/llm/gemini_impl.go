package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type geminiImpl struct {
	BaseURL string
	APIKey  string
	Model   string
}

func NewGemini(apiKey string) Llm {
	return &geminiImpl{
		BaseURL: "https://generativelanguage.googleapis.com/v1beta/models",
		APIKey:  apiKey,
		Model:   "gemini-2.5-flash-lite",
	}
}

// ----------- Request / Response Structs ------------

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// ---------------------------------------------------

func (g *geminiImpl) GetResponse(prompt string) (result ChatResponse, err error) {

	// Build request payload
	reqBody := GeminiRequest{
		Contents: []GeminiContent{
			{Parts: []GeminiPart{{Text: prompt}}},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	// Create POST request
	url := fmt.Sprintf("%s/%s:generateContent?key=%s", g.BaseURL, g.Model, g.APIKey)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var geminiResponse GeminiResponse
	if err = json.Unmarshal(body, &geminiResponse); err != nil {
		return
	}

	if len(geminiResponse.Candidates) == 0 ||
		len(geminiResponse.Candidates[0].Content.Parts) == 0 {
		err = fmt.Errorf("no response from model")
		return
	}

	result.Choices = make([]struct {
		Index   int "json:\"index\""
		Message struct {
			Role    string "json:\"role\""
			Content string "json:\"content\""
		} "json:\"message\""
		FinishReason string "json:\"finish_reason\""
	}, 1)

	result.Choices[0].Index = 1
	result.Model = g.Model
	result.Created = time.Now().Unix()
	result.ID = "gemini"
	result.Choices[0].Message.Content = geminiResponse.Candidates[0].Content.Parts[0].Text
	return
}
