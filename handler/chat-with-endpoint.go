package handler

import (
	"context"
	"encoding/json"
	"go-playground/app_config"
	"go-playground/chatgpt"
	"go-playground/redis_client"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ChatWithEndpoint(c echo.Context) error {
	topic := c.QueryParam("topic")
	chatId := c.QueryParam("chatId")
	id := c.Get("uid").(string)

	new_message := chatgpt.ChatCompletionMessage{
		Role:    chatgpt.RoleSystem,
		Content: "Bạn là một trợ lý ảo thông minh." + app_config.SYSTEM_TRAIN,
	}

	if len(topic) > 0 {
		new_message = chatgpt.ChatCompletionMessage{
			Role:    chatgpt.RoleSystem,
			Content: "Bạn là một AI hỗ trợ về " + topic + "." + app_config.SYSTEM_TRAIN,
		}
	}

	messagesStr, err := json.Marshal([]chatgpt.ChatCompletionMessage{new_message})
	if err != nil {
		return c.String(http.StatusBadRequest, "ERROR")
	}

	redisClient := redis_client.GetRedisClient()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if redisClient.Set(ctx, id+"_"+chatId, messagesStr, time.Duration(time.Second*600)).Err() != nil {
		return c.String(http.StatusBadRequest, "ERROR")
	}

	return c.String(http.StatusOK, "OK")
}
