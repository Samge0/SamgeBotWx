package u_openai

import (
	"SamgeWxApi/cmd/utils/u_http"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"strings"
)

// OpenAiParams apoenai的api请求参数
type OpenAiParams struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float32 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float32 `json:"top_p"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
}

type OpenAiResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
	Usage   Usage     `json:"usage"`
}

type Choices struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

const (
	ModelChat string = "text-davinci-003"
)

const (
	// ApiUrl api请求地址
	ApiUrl = "https://api.openai.com/v1/completions"
	// Authorization 认证token
	Authorization = "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"
	// ChatMaxTokens 回答的最大字符长度
	ChatMaxTokens = 100
)

// GetChatResponseWithToken 获取一个聊天的回答
func GetChatResponseWithToken(prompt string, maxLen int, token string) (string, error) {
	openaiParams := OpenAiParams{
		Model:            ModelChat,
		Prompt:           prompt,
		Temperature:      0.9,
		MaxTokens:        maxLen,
		TopP:             1.0,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	body, err := u_http.Post(ApiUrl, headers, openaiParams)
	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") {
			return "哎呀呀~你的问题太高深，我一时反应不过来呀", err
		} else {
			return err.Error(), err
		}
	}

	result, err := GetChatResult(*body)
	if err != nil {
		return "", err
	}
	return result, nil
}

// GetChatResponse 获取一个聊天的回答
func GetChatResponse(prompt string, maxLen int) (string, error) {
	return GetChatResponseWithToken(prompt, maxLen, Authorization)
}

// GetChatResult 获取聊天的结果
func GetChatResult(body []byte) (string, error) {
	fmt.Println(string(body))
	// 将 json 转换为结构体
	var result OpenAiResponse
	err := json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	if len(result.Choices) > 0 {
		var msg string
		msg = result.Choices[0].Text
		msg = strings.Replace(msg, "&#xA;", "", -1)
		msg = strings.Replace(msg, "\n", "", -1)
		log.Printf("回复内容：%s\n", msg)
		return msg, nil
	}
	return "", errors.New("openai Choices size is zero")
}
