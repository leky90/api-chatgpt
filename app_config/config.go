package app_config

const (
	APP_VERSION     = "v0.2"
	CHATGPT_API_KEY = "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"
	CHATGPT_MODEL   = "gpt-4"
	SYSTEM_TRAIN    = "- Trả lời đúng chuyên môn cho người đặt câu hỏi. - Tập trung trả lời câu hỏi gần nhất, các câu trước chỉ để tham khảo. Nếu chủ đề hoàn toàn khác có thể bỏ qua."
)

var (
	ChatGPTApiKey    = CHATGPT_API_KEY
	ChatGPTModel     = CHATGPT_MODEL
	CorsValidDomains = []string{"https://chatgpt.ldktech.com", "https://ldktech.com", "http://localhost:4200"}
)
