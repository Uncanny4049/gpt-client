package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

// PostConversationTalk /api/conversation/talk
// 向ChatGPT提问，等待其回复
type PostConversationTalk struct {
	Api
	Rq PostConversationTalkRQ
	Rs *PostConversationTalkRS
}

// PostConversationTalkRQ Post /api/conversation/talk
type PostConversationTalkRQ struct {
	Question        string `json:"prompt"`
	Model           string `json:"model"`
	MessageId       string `json:"message_id"`
	ParentMessageId string `json:"parent_message_id"`
	ConversationId  string `json:"conversation_id,omitempty"`
	Stream          bool   `json:"stream" default:"false"`
}

// PostConversationTalkRS Post /api/conversation/talk
type PostConversationTalkRS struct {
	ConversationId string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
	Message        struct {
		Author struct {
			Metadata struct {
			} `json:"metadata"`
			Name interface{} `json:"name"`
			Role string      `json:"role"`
		} `json:"author"`
		Content struct {
			ContentType string   `json:"content_type"`
			Parts       []string `json:"parts"`
		} `json:"content"`
		CreateTime float64 `json:"create_time"`
		EndTurn    bool    `json:"end_turn"`
		Id         string  `json:"id"`
		Metadata   struct {
			FinishDetails struct {
				StopTokens []int  `json:"stop_tokens"`
				Type       string `json:"type"`
			} `json:"finish_details"`
			IsComplete  bool   `json:"is_complete"`
			MessageType string `json:"message_type"`
			ModelSlug   string `json:"model_slug"`
			ParentId    string `json:"parent_id"`
		} `json:"metadata"`
		Recipient  string      `json:"recipient"`
		Status     string      `json:"status"`
		UpdateTime interface{} `json:"update_time"`
		Weight     float64     `json:"weight"`
	} `json:"message"`
}

func (a *PostConversationTalk) Send() {
	data, _ := json.Marshal(a.Rq)
	request, _ := http.NewRequest(a.Method, a.URL, bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	_ = json.Unmarshal(body, &(a.Rs))
}

func (a *PostConversationTalk) Default() {
	a.Name = "Talk"
	a.Method = "POST"
	a.URL = BaseUrl + "/api/conversation/talk"
}

func DefaultPostConversationTalkRQ(Question string) (a PostConversationTalkRQ) {
	a = PostConversationTalkRQ{
		Question:        Question,
		Model:           "text-davinci-002-render-sha",
		MessageId:       uuid.NewString(),
		ParentMessageId: uuid.NewString(),
		Stream:          false,
	}
	return a
}
