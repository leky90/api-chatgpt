package chatgpt

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-playground/app_config"
	"net/http"
	"strconv"
	"strings"
)

type ChanStreamResponse struct {
	Event string
	Data  string
	Error error
}
type ChatGPTClient struct {
	ApiToken string
}

const URL = "https://api.openai.com/v1/chat/completions"

func (c *ChatGPTClient) ChatCompletion(req ChatCompletionRequest, ch chan ChanStreamResponse, ctx context.Context) {

	reqJsonValue, _ := json.Marshal(req)

	httpRequest, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqJsonValue))

	if err != nil {
		fmt.Println("Create HttpRequest Error:", err)
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", "Bearer "+app_config.ChatGPTApiKey)

	// Gửi yêu cầu và nhận resp
	client := &http.Client{}

	resp, err := client.Do(httpRequest)

	if err != nil {
		fmt.Println("Call ChatGPT Error:", err)
		ch <- ChanStreamResponse{Event: "", Data: "", Error: err}
		return
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode: ", resp.StatusCode)

	if resp.StatusCode != 200 {
		ch <- ChanStreamResponse{Event: "", Data: "", Error: errors.New(strconv.Itoa(resp.StatusCode) + ":" + resp.Status)}
		return
	}

	// Đọc dữ liệu từ kết nối stream
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			// context bị hủy, dừng stream]
			return
		default:
			// Kiểm tra lỗi scanner
			if err := scanner.Err(); err != nil {
				ch <- ChanStreamResponse{Event: "", Data: "", Error: err}
				return
			}
			// Kiểm tra xem dòng dữ liệu là một sự kiện mới hay một trường dữ liệu trong sự kiện
			line := scanner.Text()
			if len(line) == 0 {
				continue
			} else if line[0] == ':' { // Giữa một số đầu dòng prng
				continue
			} else if line[0] == ' ' { // Chunk mới trong một sự kiện
				continue
			}

			event := ""
			data := ""
			parts := strings.SplitN(line, ": ", 2)

			if len(parts) == 2 {
				event = parts[0]
				data = parts[1]
				ch <- ChanStreamResponse{Event: event, Data: data, Error: nil}
			}
		}
	}
}
