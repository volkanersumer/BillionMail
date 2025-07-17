package askai

import (
	"billionmail-core/internal/service/public"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	tiktoken "github.com/pkoukk/tiktoken-go"
	openai "github.com/sashabaranov/go-openai"
)

const (
	TOKENS_CACHE_DIR = "../conf/tokens_cache" // Directory for token cache files
)

type ResultJson struct {
	ThinkStart    bool   `json:"think_start"` // Indicates if the chat is in thinking state
	ThinkEnd      bool   `json:"think_end"`   // Indicates if the chat has ended thinking state
	IsThink       bool   `json:"is_think"`    // Indicates if the chat is in thinking state
	Content       string `json:"content"`     // Content of the chat message
	Reasoning     string `json:"reasoning"`   // Reasoning or explanation for the chat message
	IsHtmlContent bool   `json:"is_html"`     // HTML content of the chat message
	Role          string `json:"role"`        // Role of the message sender (e.g., "user", "assistant")
	IsEnd         bool   `json:"is_end"`      // Indicates if this is the end of the chat message
	UpdateTime    int64  `json:"update_time"` // Last update time of the chat message
}

type OpenAI struct {
	ApiKey                  string         // API key for OpenAI
	BaseUrl                 string         // Base URL for OpenAI API
	ModelId                 string         // Model ID for OpenAI API
	ChatId                  string         // Chat ID for OpenAI API
	SupplierName            string         // Supplier name for OpenAI
	MaxTokens               int            // Maximum tokens for OpenAI API
	Client                  *openai.Client // OpenAI client instance
	Ctx                     context.Context
	Reasoning               string    // Reasoning for the chat message
	ThinkStart              bool      // Indicates if the chat is in thinking state
	ThinkEnd                bool      // Indicates if the chat has ended thinking state
	IsThinking              bool      // Indicates if the chat is in thinking state
	IsThinkContent          bool      // Indicates if the chat content is being thought about
	HtmlContent             string    // HTML content of the chat message
	IsHtml                  bool      // Indicates if the content is HTML
	LastContent             string    // Last content sent in the chat
	MessageId               string    // Unique identifier for the message
	CreateTime              int64     // Creation time of the chat
	RequestTime             int64     // Time when the request was made
	Usage                   ChatUsage // Usage statistics for the chat
	IsRecordHtml            bool      // Indicates if HTML content should be recorded
	Content                 string    // Content of the chat message
	AddedSystemPrompt       string    // Initial system prompt for the chat
	AddedSystemPromptTokens int       // Tokens for the initial system prompt
	TempInfo                *TempChat // Chat information for the OpenAI chat
	UserContent             string    // User content for the chat
	IsModify                bool      // Indicates if the chat is being modified

}

func NewOpenAI(ctx context.Context, apiKey, baseUrl, modelId, chatId, supplierName string, maxTokens int) *OpenAI {
	tempInfo, err := GetTempChat(chatId)
	if err != nil {
		fmt.Println("Error getting template chat:", err)
		return nil // Return nil if there is an error getting template chat info
	}
	return &OpenAI{
		ApiKey:       apiKey,
		BaseUrl:      baseUrl,
		ModelId:      modelId,
		ChatId:       chatId,
		SupplierName: supplierName,
		MaxTokens:    maxTokens,
		Client:       nil, // Client will be initialized later
		Ctx:          ctx,
		Reasoning:    "",    // Default reasoning is empty
		IsThinking:   false, // Default thinking state is false
		HtmlContent:  "",    // Default HTML content is empty
		IsHtml:       false, // Default content type is not HTML
		MessageId:    "",
		CreateTime:   public.GetNowTime(), // Set creation time to the current time
		RequestTime:  0,                   // Request time will be set when the chat is initiated
		Usage: ChatUsage{
			PromptTokens:     0, // Default prompt tokens is 0
			CompletionTokens: 0, // Default completion tokens is 0
			TotalTokens:      0, // Default total tokens is 0
		},
		IsRecordHtml:            false,    // Default is not to record HTML content
		LastContent:             "",       // Default last content is empty
		Content:                 "",       // Default content is empty
		AddedSystemPrompt:       "",       // Set the initial system prompt
		AddedSystemPromptTokens: 0,        // Set initial system prompt tokens
		TempInfo:                tempInfo, // Get chat information based on the chat ID
		UserContent:             "",       // Default user content is empty
		IsModify:                false,    // Default is not to modify the chat
	}
}

type ModelConfig struct {
	Encoding    string
	MaxTokens   int
	Description string
}

var SYSTEM_PROMPT_TOKENS = 0

// GetClient returns the OpenAI client
func (o *OpenAI) GetClient() *openai.Client {
	if o.SupplierName == "openai" {
		o.Client = openai.NewClient(o.ApiKey)
	} else {
		config := openai.DefaultConfig(o.ApiKey)
		config.BaseURL = o.BaseUrl

		o.Client = openai.NewClientWithConfig(config)
	}
	return o.Client
}

// GetCacheTokens retrieves the cached tokens for a specific key
// It checks if the cache directory exists, creates it if not, and reads the tokens from a file named after the cache key.
func (o *OpenAI) GetCacheTokens(cacheKey string) int {
	// cache filename
	if !public.FileExists(TOKENS_CACHE_DIR) {
		os.MkdirAll(TOKENS_CACHE_DIR, 0755) // Create the cache directory if it does not exist
	}
	cacheFile := TOKENS_CACHE_DIR + "/" + cacheKey
	if !public.FileExists(cacheFile) {
		return 0
	}

	// read cache file
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		fmt.Println("Error reading cache file:", err)
		return 0 // Return 0 if there is an error reading the cache file
	}

	// string to int
	tokens, err := strconv.Atoi(string(data)) // Convert the string data to an integer
	if err != nil {
		return 0 // Return 0 if there is an error converting the string to an integer
	}
	return tokens
}

// SetCacheTokens sets the cached tokens for a specific key
// It creates a cache directory if it does not exist and writes the tokens to a file named after the cache key.
func (o *OpenAI) SetCacheTokens(cacheKey string, tokens int) error {
	// cache filename
	if !public.FileExists(TOKENS_CACHE_DIR) {
		os.MkdirAll(TOKENS_CACHE_DIR, 0755) // Create the cache directory if it does not exist
	}
	cacheFile := TOKENS_CACHE_DIR + "/" + cacheKey

	// write cache file
	err := os.WriteFile(cacheFile, []byte(strconv.Itoa(tokens)), 0644)
	if err != nil {
		fmt.Println("Error writing cache file:", err)
		return err
	}
	return nil
}

// 计算tokens
func (o *OpenAI) CalculateTokens(content string, isCache bool) (int, error) {
	// Implementation for calculating tokens based on the content
	// This function should handle the logic for calculating tokens and return the count or any error encountered

	// 尝试从缓存获取tokens
	cacheKey := public.Md5(content)
	if isCache {
		cachedTokens := o.GetCacheTokens(cacheKey)
		if cachedTokens > 0 {
			return cachedTokens, nil // Return cached tokens if available
		}
	}

	encoding := "cl100k_base"
	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		return 0, err // Return error if encoding is not found
	}
	tokens := tke.Encode(content, nil, nil)
	tokenCount := len(tokens)

	// 将计算的tokens存入缓存
	if isCache {
		o.SetCacheTokens(cacheKey, tokenCount)
	}

	return tokenCount, nil // Placeholder implementation, replace with actual token calculation logic
}

// CalculatePromptTokens calculates the number of prompt tokens based on the content length
// It also adds the system prompt tokens to the total prompt tokens
func (o *OpenAI) CalculatePromptTokens(content string) {
	o.Usage.PromptTokens, _ = o.CalculateTokens(content, false) // Update prompt tokens based on the content length
	if SYSTEM_PROMPT_TOKENS == 0 {
		SYSTEM_PROMPT_TOKENS, _ = o.CalculateTokens(INITIAL_SYSTEM_PROMPT, true)
	}
	o.Usage.PromptTokens += SYSTEM_PROMPT_TOKENS + o.AddedSystemPromptTokens // Add system prompt tokens to the total prompt tokens
}

func (o *OpenAI) WriteMessage(content string, finishReason string) {

	userMessage := Message{
		MessageId:     o.MessageId,
		Ppid:          "", // Parent message ID is empty for the initial message
		ChatId:        o.ChatId,
		Role:          openai.ChatMessageRoleUser,
		Content:       content,
		HtmlContent:   "",
		CreateTime:    o.CreateTime,        // Use the creation time set in the OpenAI struct
		EndTime:       public.GetNowTime(), // End time is 0 for the initial message
		Reasoning:     "",
		SupplierName:  o.SupplierName,
		ModelId:       o.ModelId,
		TimeConsuming: 0, // Time consuming is 0 for the initial message
		FinishReason:  "",
		Usage: ChatUsage{
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		},
	}
	userMessage.TimeConsuming = userMessage.EndTime - userMessage.CreateTime // Calculate time consuming

	assistantMessage := Message{
		MessageId:     GetUUID(),
		Ppid:          o.MessageId, // Use the MessageId as the parent message ID
		ChatId:        o.ChatId,
		Role:          openai.ChatMessageRoleAssistant,
		Content:       o.Content,
		HtmlContent:   o.HtmlContent,
		CreateTime:    o.RequestTime,
		EndTime:       public.GetNowTime(), // End time is set to the current time
		Reasoning:     o.Reasoning,
		SupplierName:  o.SupplierName,
		ModelId:       o.ModelId,
		TimeConsuming: 0, // Time consuming is 0 for the initial message
		FinishReason:  finishReason,
		Usage: ChatUsage{
			PromptTokens:     o.Usage.PromptTokens,
			CompletionTokens: o.Usage.CompletionTokens,
			TotalTokens:      o.Usage.PromptTokens + o.Usage.CompletionTokens,
		},
	}
	assistantMessage.TimeConsuming = assistantMessage.EndTime - assistantMessage.CreateTime // Calculate time consuming

	o.AppendMessage(userMessage, assistantMessage) // Append the messages to the chat
}

func (o *OpenAI) GetCompanyProfilePrompt() (string, int) {
	companyProFile, err := GetCompanyProfile(o.TempInfo.Domain) // Get the company profile for the chat
	if err != nil {
		fmt.Println("Error getting company profile:", err)
		return "", 0
	}
	companyProFileJson, err := json.Marshal(companyProFile) // Marshal the company profile to JSON
	if err != nil {
		fmt.Println("Error marshalling company profile:", err)
		return "", 0
	}

	prompt := "Company Profile: " + string(companyProFileJson) + "\n\n" // Return the company profile as a string
	tokens, _ := o.CalculateTokens(prompt, true)                        // Calculate the tokens for the company profile prompt

	return prompt, tokens
}

func (o *OpenAI) GetProjectConfigPrompt() (string, int) {
	projectConfig, err := ReadProjectConfig(o.TempInfo.Domain) // Get the project configuration for the chat
	if err != nil {
		fmt.Println("Error reading project config:", err)
		return "", 0
	}
	projectConfigJson, err := json.Marshal(projectConfig) // Marshal the project configuration to JSON
	if err != nil {
		fmt.Println("Error marshalling project config:", err)
		return "", 0
	}
	prompt := "Project Config: " + string(projectConfigJson) + "\n\n" // Return the project configuration as a string
	tokens, _ := o.CalculateTokens(prompt, true)                      // Calculate the tokens for the project configuration prompt
	return prompt, tokens
}

func (o *OpenAI) GetSitemapPrompt() (string, int) {
	sitemap, err := GetSiteMap(o.TempInfo.Domain) // Get the sitemap for the chat
	if err != nil {
		fmt.Println("Error getting sitemap:", err)
		return "", 0
	}

	sitemapJson, err := json.Marshal(sitemap) // Marshal the sitemap to JSON
	if err != nil {
		fmt.Println("Error marshalling sitemap:", err)
		return "", 0
	}
	prompt := "Sitemap: " + string(sitemapJson) + "\n\n" // Return the sitemap as a string
	tokens, _ := o.CalculateTokens(prompt, true)         // Calculate the tokens for the sitemap prompt
	return prompt, tokens
}

func (o *OpenAI) GetStylePrompt() (string, int) {
	styleConfig, err := GetStyleConfig(o.TempInfo.Domain) // Get the style configuration for the chat
	if err != nil {
		fmt.Println("Error getting style config:", err)
		return "", 0
	}
	styleConfigJson, err := json.Marshal(styleConfig) // Marshal the style configuration to JSON
	if err != nil {
		fmt.Println("Error marshalling style config:", err)
		return "", 0
	}
	prompt := "Style Config: " + string(styleConfigJson) + "\n\n" // Return the style configuration as a string
	tokens, _ := o.CalculateTokens(prompt, true)                  // Calculate the tokens for the style configuration prompt
	return prompt, tokens
}

func (o *OpenAI) GetFooterPrompt() (string, int) {
	footer, err := GetFooter(o.TempInfo.Domain) // Get the footer for the chat
	if err != nil {
		fmt.Println("Error getting footer:", err)
		return "", 0
	}
	footerJson, err := json.Marshal(footer) // Marshal the footer to JSON
	if err != nil {
		fmt.Println("Error marshalling footer:", err)
		return "", 0
	}
	prompt := "Footer: " + string(footerJson) + "\n\n" // Return the footer as a string
	tokens, _ := o.CalculateTokens(prompt, true)       // Calculate the tokens for the footer prompt
	return prompt, tokens
}

func (o *OpenAI) GetKnowledgeBasePrompt() (string, int) {
	knowledgeBase, err := GetKnowledgeBaseList(o.TempInfo.Domain) // Get the knowledge base for the chat
	if err != nil {
		fmt.Println("Error getting knowledge base:", err)
		return "", 0
	}

	// 取内容长度小于1024的文档
	var filteredKnowledgeBase []string
	for i := len(knowledgeBase) - 1; i >= 0; i-- {
		if len(knowledgeBase[i].Content) < 1024 {
			filteredKnowledgeBase = append(filteredKnowledgeBase, knowledgeBase[i].Content)
		}
	}

	knowledgeBaseJson, err := json.Marshal(filteredKnowledgeBase) // Marshal the knowledge base to JSON
	if err != nil {
		fmt.Println("Error marshalling knowledge base:", err)
		return "", 0
	}
	prompt := "Knowledge Base Doc List: " + string(knowledgeBaseJson) + "\n\n" // Return the knowledge base as a string
	tokens, _ := o.CalculateTokens(prompt, true)                               // Calculate the tokens for the knowledge base prompt
	return prompt, tokens
}

func (o *OpenAI) GetImagesPrompt() (string, int) {
	images, err := ReadImagesConfig(o.TempInfo.Domain) // Get the images configuration for the chat
	if err != nil {
		fmt.Println("Error reading images list:", err)
		return "", 0
	}
	if err != nil {
		fmt.Println("Error reading images list:", err)
		return "", 0
	}
	imagesJson, err := json.Marshal(images) // Marshal the images configuration to JSON
	if err != nil {
		fmt.Println("Error marshalling images list:", err)
		return "", 0
	}
	prompt := "Images: " + string(imagesJson) + "\n\n" // Return the images configuration as a string
	tokens, _ := o.CalculateTokens(prompt, true)       // Calculate the tokens for the images prompt
	return prompt, tokens
}

func (o *OpenAI) GetLastPrompt() (string, int) {
	prompt := "The following are reference materials, all in JSON format, which you can use as needed:\n\n"
	tokens, _ := o.CalculateTokens(prompt, true) // Calculate the tokens for the last prompt
	return prompt, tokens
}

func (o *OpenAI) GetSystemPrompt() string {
	if o.AddedSystemPrompt == "" {

		o.AddedSystemPrompt = ""
		prompt := ""
		tokens := 0
		prompt, tokens = o.GetLastPrompt() // Get the last prompt
		o.AddedSystemPrompt += prompt      // Add the last prompt to the system prompt
		o.AddedSystemPromptTokens = tokens // Set the tokens for the last prompt
		// Get the additional system prompts
		prompt, tokens = o.GetCompanyProfilePrompt() // Get the company profile prompt
		o.AddedSystemPrompt += prompt                // Add the company profile prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the company profile prompt
		prompt, tokens = o.GetProjectConfigPrompt()  // Get the project configuration prompt
		o.AddedSystemPrompt += prompt                // Add the project configuration prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the project configuration prompt
		prompt, tokens = o.GetSitemapPrompt()        // Get the sitemap prompt
		o.AddedSystemPrompt += prompt                // Add the sitemap prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the sitemap prompt
		prompt, tokens = o.GetStylePrompt()          // Get the style prompt
		o.AddedSystemPrompt += prompt                // Add the style prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the style prompt
		prompt, tokens = o.GetFooterPrompt()         // Get the footer prompt
		o.AddedSystemPrompt += prompt                // Add the footer prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the footer prompt
		prompt, tokens = o.GetKnowledgeBasePrompt()  // Get the knowledge base prompt
		o.AddedSystemPrompt += prompt                // Add the knowledge base prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the knowledge base prompt
		prompt, tokens = o.GetImagesPrompt()         // Get the images prompt
		o.AddedSystemPrompt += prompt                // Add the images prompt to the system prompt
		o.AddedSystemPromptTokens += tokens          // Add the tokens for the images prompt

		o.AddedSystemPrompt += fmt.Sprintf("Current time: %s\n\n", public.GetNowDate())
	}
	systemPrompt := INITIAL_SYSTEM_PROMPT
	if o.IsModify {
		systemPrompt = FOLLOW_UP_SYSTEM_PROMPT
	}

	return systemPrompt + "\n\n" + o.AddedSystemPrompt // Return the initial system prompt concatenated with any additional system prompt

}

// GetMessages retrieves the messages for the chat
// It reads the chat history, adds the system prompt, and appends the user content to the messages.
// It also handles the last assistant message and ensures the format complies with OpenAI requirements.
func (o *OpenAI) GetMessages() []openai.ChatCompletionMessage {
	historyMessages := o.ReadMessages() // Read messages from the chat history
	lastContent, _ := GetHtml(o.ChatId)
	o.IsModify = false
	if len(historyMessages) > 0 && lastContent != "" {
		o.IsModify = true // If there are history messages and last content is not empty, set IsModify to true
	}

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: o.GetSystemPrompt()},
	}

	// Process history messages and ensure proper alternating pattern
	for _, message := range historyMessages {
		// Truncate long assistant messages
		content := message.Content
		if message.Role == openai.ChatMessageRoleAssistant && len(content) > 512 {
			content = "..."
		}

		messages = append(messages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: content,
		})
	}

	// Ensure the conversation ends with a user message for proper OpenAI format
	messagesLen := len(messages)
	if messagesLen > 1 && messages[messagesLen-1].Role == openai.ChatMessageRoleAssistant {
		// If last message is from assistant, update its content with HTML if available
		if lastContent != "" {
			messages[messagesLen-1].Content = lastContent
		}
	}

	// Add the current user content to the messages
	if strings.TrimSpace(o.UserContent) != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: o.UserContent,
		})
	}

	return messages
}

func (o *OpenAI) Chat(content string, isText bool) error {
	// Implementation for sending a chat message to OpenAI
	// This function should handle the logic for sending a chat message and return the response or any error encountered

	request := ghttp.RequestFromCtx(o.Ctx)
	if request == nil {
		return errors.New("request context is nil")
	}

	if o.Client == nil {
		o.GetClient()
	}

	o.UserContent = content // Set the user content to the provided content

	req := openai.ChatCompletionRequest{
		Model:               o.ModelId,
		MaxCompletionTokens: o.MaxTokens,
		Messages:            o.GetMessages(), // Get the messages for the chat
		Temperature:         0.6,             // Set temperature for the chat completion
		Stream:              true,
	}
	o.MessageId = GetUUID()

	o.RequestTime = public.GetNowTime() // Set request time to the current time
	o.CalculatePromptTokens(content)    // Calculate prompt tokens based on the content

	stream, err := o.Client.CreateChatCompletionStream(o.Ctx, req)
	if err != nil {
		return err
	}

	defer stream.Close()

	// set request.Response.Header().Set("Content-Type", "text/event-stream")
	request.Response.Header().Set("Content-Type", "text/event-stream; charset=utf-8") // Set content type for SSE (Server-Sent Events)
	request.Response.Header().Set("Cache-Control", "no-cache")
	request.Response.Header().Set("Connection", "keep-alive")
	request.Response.WriteHeader(200) // OK status
	finishReason := ""
	for {
		if ChatStatus[o.ChatId] == false {
			finishReason = "stop"
			break
		}

		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		// Set Usage.PromptTokens, CompletionTokens, and TotalTokens
		o.Usage.CompletionTokens += len(response.Choices[0].Delta.Content)
		o.Usage.TotalTokens += o.Usage.PromptTokens + o.Usage.CompletionTokens
		finishReason = string(response.Choices[0].FinishReason)

		messageBody := response.Choices[0].Delta.Content
		if messageBody == "" {
			if response.Choices[0].Delta.ReasoningContent != "" {
				if !o.IsThinkContent {
					o.ThinkStart = true
				} else {
					o.ThinkStart = false
				}
				o.IsThinking = true
				o.IsThinkContent = true                                  // Set thinking content state to true
				messageBody = response.Choices[0].Delta.ReasoningContent // Use reasoning content
			} else {
				continue // Skip empty messages
			}
		} else {
			if o.IsThinkContent {
				o.ThinkEnd = true                          // Set thinking end state to true if content is not empty
				o.WriteEvent(request, "\n", isText, false) // Write event to the response
				o.ThinkEnd = false
				o.IsThinkContent = false // Reset thinking content state after processing
				if o.IsThinking {
					o.IsThinking = false // Reset thinking state after processing
				}
			}
		}

		// 获取HTML内容，格式：```html
		// if messageBody

		o.WriteEvent(request, messageBody, isText, false) // Write event to the response
	}
	// Ensure the response is properly closed
	o.WriteHtml()                           // Write HTML content if the chat has ended
	o.WriteMessage(content, finishReason)   // Write the final message to the chat
	o.WriteEvent(request, "", isText, true) // Write final event to indicate completion
	return nil
}

type ReplaceHtml struct {
	SearchContent  string // Content to search for in the HTML
	ReplaceContent string // Content to replace the searched content with
}

func (o *OpenAI) ReplaceHtml() {

	if !o.IsModify {
		return // If not modifying, return early
	}

	// If modifying, replace the HTML content with the formatted content
	if o.Content == "" {
		return
	}

	replaceHtml := []ReplaceHtml{}
	if strings.Contains(o.Content, SEARCH_START) && strings.Contains(o.Content, REPLACE_END) && strings.Contains(o.Content, DIVIDER) {

		// Split the content into parts based on the SEARCH_START and REPLACE_END markers
		parts := strings.Split(o.Content, SEARCH_START)
		for _, part := range parts {
			if strings.Contains(part, REPLACE_END) {
				// Further split each part to get the search and replace content
				subParts := strings.Split(part, REPLACE_END)
				if len(subParts) == 2 && strings.Contains(subParts[0], DIVIDER) {
					searchReplace := strings.Split(subParts[0], DIVIDER)
					if len(searchReplace) == 2 {
						replaceHtml = append(replaceHtml, ReplaceHtml{
							SearchContent:  searchReplace[0],
							ReplaceContent: searchReplace[1],
						})
					}
				}
			}
		}
	}

	htmlContent, err := GetHtml(o.ChatId) // Get the HTML content for the chat
	if err != nil {
		fmt.Println("Error getting HTML content:", err)
		return // Return if there is an error getting the HTML content
	}

	// Replace the HTML content based on the search and replace rules
	for _, replace := range replaceHtml {
		if replace.SearchContent != "" && replace.ReplaceContent != "" {
			htmlContent = strings.ReplaceAll(htmlContent, replace.SearchContent, replace.ReplaceContent)
		}
	}

	err = ModifyHtml(o.ChatId, htmlContent) // Modify the HTML content for the chat
	if err != nil {
		fmt.Println("Error modifying HTML content:", err)
	}

	return

}

// WriteHtml writes the HTML content to a file if it exists
func (o *OpenAI) WriteHtml() {
	if o.IsModify {
		o.ReplaceHtml() // Replace HTML content if modifying
		return
	}
	if o.HtmlContent != "" {
		filename := CHAT_CONFIG_PATH + "/" + o.ChatId + "/code.html"
		o.HtmlContent = strings.Trim(o.HtmlContent, "`")
		if strings.HasPrefix(o.HtmlContent, "html") {
			o.HtmlContent = o.HtmlContent[4:] // Remove the "html" prefix if present
		}
		o.HtmlContent = strings.TrimSpace(o.HtmlContent) // Trim leading and trailing whitespace
		o.HtmlContent = strings.Trim(o.HtmlContent, "`")
		os.WriteFile(filename, []byte(o.HtmlContent), os.ModePerm)
		g.DB().Model("email_templates").Where("chat_id", o.ChatId).Update(g.Map{
			"content": o.HtmlContent, // Update the HTML content in the database
		})
	}
}

func (o *OpenAI) ReadMessages() []Message {
	messageFile := CHAT_CONFIG_PATH + "/" + o.ChatId + "/message.json"
	var messages []Message
	if !public.FileExists(messageFile) {
		return messages // Return nil if the message file does not exist
	}

	// Read existing messages from the file
	fileContent, err := os.ReadFile(messageFile)
	if err != nil {
		fmt.Println("Failed to read message file: ", err)
		return messages
	}

	err = json.Unmarshal(fileContent, &messages)
	if err != nil {
		fmt.Println("Failed to unmarshal messages: ", err)
		return messages
	}

	return messages
}

// SaveMessages saves the messages to a file
// It creates a directory for the chat if it does not exist and writes the messages to a JSON file
func (o *OpenAI) SaveMessages(messages []Message) error {
	messageFile := CHAT_CONFIG_PATH + "/" + o.ChatId + "/message.json"
	updatedContent, err := json.Marshal(messages)
	if err != nil {
		return err
	}
	return os.WriteFile(messageFile, updatedContent, os.ModePerm)
}

func (o *OpenAI) AppendMessage(message ...Message) {
	messages := o.ReadMessages() // Read existing messages from the file

	// Append the new message to the existing messages
	messages = append(messages, message...)

	// Save the updated messages back to the file
	err := o.SaveMessages(messages)
	if err != nil {
		fmt.Println("Failed to save messages: ", err) // Log error if saving messages fails
		return
	}
}

// SetHtmlContent sets the HTML content and reasoning for the OpenAI chat
func (o *OpenAI) SetHtmlContent(content string, isText bool) (reasoning string, IsHtmlContent bool) {
	reasoning = ""
	IsHtmlContent = false
	trimSpaceContent := strings.TrimSpace(content)

	// Check for special tags and update the state accordingly
	if trimSpaceContent == "<think>" {
		o.IsThinking = true
	} else if trimSpaceContent == "</think>" {
		o.IsThinking = false
		if o.Reasoning != "" {
			// If reasoning is not empty, append the content to reasoning
			reasoning = content[:]
			if !isText {
				content = ""
			}
		}
		o.Reasoning += content
	}

	// Check for HTML code blocks and toggle the HTML state
	if trimSpaceContent == "```" || (trimSpaceContent == "`" && o.LastContent == "``") {
		o.IsHtml = !o.IsHtml // Toggle HTML state
		if !o.IsHtml {
			IsHtmlContent = true
			if !o.IsRecordHtml {
				o.HtmlContent += content
				o.IsRecordHtml = true
			}
		}
	}

	// isModify is true, it means the user is modifying the chat
	// if o.IsModify {

	// }

	// If the content is a thinking block, append it to the Reasoning
	if o.IsThinking {
		o.Reasoning += content
		reasoning = content[:]
		if !isText {
			content = ""
		}
	}

	// If the content is HTML, append it to the HtmlContent
	if o.IsHtml {
		if !o.IsRecordHtml {
			o.HtmlContent += content
		}
		IsHtmlContent = true
	}

	o.LastContent = content // Store the last content sent in the chat

	return reasoning, IsHtmlContent
}

func (o *OpenAI) WriteEvent(request *ghttp.Request, content string, isText bool, isEnd bool) error {
	// Implementation for writing an event to the OpenAI response
	// This function should handle the logic for writing an event and return any error encountered
	Write := request.Response.Write
	if isEnd {
		Write = request.Response.WriteExit
	} else {
		o.Usage.CompletionTokens += 1 // Increment completion tokens for the final message
	}

	reasoning, IsHtmlContent := o.SetHtmlContent(content, isText) // Set HTML content and reasoning
	o.Content += content                                          // Update the content in the OpenAI struct
	if isText {
		Write(content)
	} else {
		resultJson := ResultJson{
			Content:       content,
			Reasoning:     reasoning,
			IsHtmlContent: IsHtmlContent,
			Role:          openai.ChatMessageRoleAssistant,
			IsEnd:         isEnd,
			ThinkStart:    o.ThinkStart,
			ThinkEnd:      o.ThinkEnd,
			IsThink:       o.IsThinking,
			UpdateTime:    public.GetNowTime(),
		}
		strJson, err := json.Marshal(resultJson)
		if err != nil {
			return err // Handle JSON marshaling error
		}
		Write("id: " + o.ChatId + "\ndata: " + string(strJson) + "\n\n")
	}

	request.Response.Flush()
	return nil
}
