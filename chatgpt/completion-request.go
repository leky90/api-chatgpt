package chatgpt

const RoleUser = "user"
const RoleAssistant = "assistant"
const RoleSystem = "system"

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Định nghĩa kiểu dữ liệu request
type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	MaxTokens        int                     `json:"max_tokens,omitempty"`
	Temperature      float32                 `json:"temperature,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"`
	N                interface{}             `json:"n,omitempty"`
	Stop             int                     `json:"stop,omitempty"`
	User             string                  `json:"user,omitempty"`
	Stream           bool                    `json:"stream,omitempty"`
}
