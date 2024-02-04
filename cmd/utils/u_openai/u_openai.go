package u_openai

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/utils/u_http"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"strings"
)

// OpenAiParams apoenai的api请求参数
type OpenAiParams struct {
	Model            string     `json:"model"`
	Temperature      float32    `json:"temperature"`
	MaxTokens        int        `json:"max_tokens"`
	TopP             float32    `json:"top_p"`
	FrequencyPenalty float32    `json:"frequency_penalty"`
	PresencePenalty  float32    `json:"presence_penalty"`
	Messages         []Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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
	Message      Messages `json:"message"`
	FinishReason string   `json:"finish_reason"`
	Index        int      `json:"index"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// GetChatResponseWithToken 获取一个聊天的回答
func GetChatResponseWithToken(prompt string) (string, error) {
	message := Messages{
		Role:    "user",
		Content: prompt,
	}
	openaiParams := OpenAiParams{
		Model:            config.LoadConfig().Model,
		Temperature:      config.LoadConfig().Temperature,
		MaxTokens:        config.LoadConfig().MaxTokens,
		TopP:             config.LoadConfig().TopP,
		FrequencyPenalty: config.LoadConfig().FrequencyPenalty,
		PresencePenalty:  config.LoadConfig().PresencePenalty,
		Messages:         []Messages{message},
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", config.LoadConfig().ApiKey),
	}
	reqUrl := fmt.Sprintf("%s/chat/completions", config.LoadConfig().BaseUrl)
	body, err := u_http.Post(reqUrl, headers, openaiParams)
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
		msg = result.Choices[0].Message.Content
		msg = strings.Replace(msg, "&#xA;", "", -1)
		msg = strings.Replace(msg, "\n", "", -1)
		log.Printf("回复内容：%s\n", msg)
		return msg, nil
	}
	return "", errors.New("openai Choices size is zero")
}
