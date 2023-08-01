package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"go-playground/app_config"
	"go-playground/chatgpt"
	"go-playground/redis_client"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ChatEndpoint(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Api-version", app_config.APP_VERSION)
	id := c.Get("uid").(string)
	message := c.QueryParam("message")
	retry := c.QueryParam("retry")
	chatId := c.QueryParam("chatId")

	new_message := chatgpt.ChatCompletionMessage{
		Role:    chatgpt.RoleUser,
		Content: message,
	}

	redisClient := redis_client.GetRedisClient()

	strMessages := redisClient.Get(context.Background(), id+"_"+chatId).Val()

	var messages []chatgpt.ChatCompletionMessage

	if len(strMessages) > 0 {
		err := json.Unmarshal([]byte(strMessages), &messages)
		if err != nil {
			fmt.Println("=============ERROR=============")
			fmt.Printf("Parse Json error: %v\n", err)

			return c.String(http.StatusBadRequest, "Error")
		}
	}

	if retry != "true" || len(messages) == 0 {
		if len(messages) > 0 {
			messages = append(messages, new_message)
		} else {
			system_message := chatgpt.ChatCompletionMessage{
				Role:    chatgpt.RoleSystem,
				Content: "Bạn là một trợ lí AI thông minh và hữu ích." + app_config.SYSTEM_TRAIN,
			}
			messages = []chatgpt.ChatCompletionMessage{system_message, new_message}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	messagesToStr, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("=============ERROR=============")
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "ERROR")
	}

	if err = redisClient.Set(ctx, id+"_"+chatId, messagesToStr, time.Duration(time.Second*600)).Err(); err != nil {
		fmt.Println("=============ERROR=============")
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "ERROR")
	}

	fmt.Println("=============START=============")
	fmt.Println("Stream start: ", time.Now().Format("15:04:05"), len(messages), cap(messages))

	client := &chatgpt.ChatGPTClient{
		ApiToken: app_config.ChatGPTApiKey,
	}

	fmt.Println("====USING " + app_config.ChatGPTModel + " MODEL====")
	chatCompletionRequest := chatgpt.ChatCompletionRequest{
		Model:    app_config.ChatGPTModel,
		Messages: messages,
		User:     id,
		Stream:   true,
	}

	responseChan := make(chan chatgpt.ChanStreamResponse)
	defer close(responseChan)

	// Gửi câu hỏi tới Chatbot và hiển thị kết quả trả lời
	go client.ChatCompletion(chatCompletionRequest, responseChan, ctx)

	enc := json.NewEncoder(c.Response())

	for response := range responseChan {
		if err := enc.Encode(response); err != nil {
			fmt.Println("=============ERROR=============")
			fmt.Println(err)
			break
		}

		if response.Error != nil {
			fmt.Println("=============ERROR=============")
			fmt.Println(response.Error)
			break
		}

		if response.Data == "[DONE]" {
			fmt.Fprintf(c.Response().Writer, "event: done\n\n")
			fmt.Println("=============END=============")
			return nil
		}

		fmt.Fprintf(c.Response().Writer, "data: %s\n\n", response.Data)

		c.Response().Flush()
	}

	return c.String(http.StatusBadRequest, "Error")
}
