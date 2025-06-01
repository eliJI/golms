package golms

type Requester struct {
	Model string
}

type Response struct {
	Id        string    `json:"id"`
	Object    string    `json:"object"`
	Created   int       `json:"created"`
	Model     string    `json:"model"`
	Choices   []Choice  `json:"choices"`
	Usage     Usage     `json:"usage"`
	Stats     Stats     `json:"stats"`
	ModelInfo ModelInfo `json:"model_info"`
	Runtime   Runtime   `json:"runtime"`
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Choice struct {
	Index        int     `json:"index"`
	Logprobs     any     `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
type Stats struct {
	TokensPerSecond  float64 `json:"tokens_per_second"`
	TimeToFirstToken float64 `json:"time_to_first_token"`
	GenerationTime   float64 `json:"generation_time"`
	StopReason       string  `json:"stop_reason"`
}
type ModelInfo struct {
	Arch          string `json:"arch"`
	Quant         string `json:"quant"`
	Format        string `json:"format"`
	ContextLength int    `json:"context_length"`
}
type Runtime struct {
	Name             string   `json:"name"`
	Version          string   `json:"version"`
	SupportedFormats []string `json:"supported_formats"`
}
