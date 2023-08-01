package chatgpt

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type MessageChoice struct {
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
	Index        int                   `json:"index"`
}

type ChatCompletionResponse struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created int             `json:"create"`
	Model   string          `json:"model"`
	Usage   Usage           `json:"usage`
	Choices []MessageChoice `json:"choices"`
}
