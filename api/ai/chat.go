package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReqChat struct {
	Chat string `json:"chat"`
}

type ResChat struct {
	Chat string `json:"chat"`
}

type ReqOpenAI struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      float64  `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             int      `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type ResOpenAI struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func Chat(c *gin.Context) {
	r := ReqChat{}

	err := c.Bind(&r)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"description": "bad params",
		})
		return
	}

	prompt := r.Chat
	input := strings.Split(prompt, "\n")
	inputCount := len(input)
	if inputCount > 20 {
		prompt = strings.Join(input[inputCount-6:], "\n")
	}

	reqOpenAI := ReqOpenAI{
		Model:            "text-davinci-002",
		Prompt:           prompt,
		Temperature:      0.5,
		MaxTokens:        100,
		TopP:             1,
		FrequencyPenalty: 0.5,
		PresencePenalty:  0.6,
		Stop:             []string{"Human", "AI"},
	}
	data, err := json.Marshal(reqOpenAI)
	if err != nil {
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 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")

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	if len(body) == 0 {
		return
	}

	resBody := ResOpenAI{}
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return
	}

	resChat := ResChat{}
	resChat.Chat = r.Chat + strings.TrimSpace(resBody.Choices[0].Text)

	c.JSON(http.StatusOK, resChat)
}
