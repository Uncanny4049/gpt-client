package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// GetConversation /api/conversation/<conversation_id>
// 通过会话ID获取指定会话详情
type GetConversation struct {
	Api
	Rq GetConversationRQ
	Rs GetConversationRS
}

// GetConversationRQ Get /api/conversation/<conversation_id>
type GetConversationRQ struct {
	ConversationId string
}

// GetConversationRS Get /api/conversation/<conversation_id>
type GetConversationRS struct {
	Title      string  `json:"title"`
	CreateTime float64 `json:"create_time"`
	UpdateTime float64 `json:"update_time"`
	Mapping    map[string]struct {
		Id      string `json:"id"`
		Message struct {
			Id     string `json:"id"`
			Author struct {
				Role     string `json:"role"`
				Metadata struct {
				} `json:"metadata"`
			} `json:"author"`
			CreateTime float64 `json:"create_time"`
			Content    struct {
				ContentType string   `json:"content_type"`
				Parts       []string `json:"parts"`
			} `json:"content"`
			Status   string  `json:"status"`
			EndTurn  bool    `json:"end_turn"`
			Weight   float64 `json:"weight"`
			Metadata struct {
				FinishDetails struct {
					Type       string `json:"type"`
					StopTokens []int  `json:"stop_tokens"`
				} `json:"finish_details"`
				IsComplete bool   `json:"is_complete"`
				ModelSlug  string `json:"model_slug"`
				ParentId   string `json:"parent_id"`
				Timestamp  string `json:"timestamp_"`
			} `json:"metadata"`
			Recipient string `json:"recipient"`
		} `json:"message"`
		Parent   string        `json:"parent"`
		Children []interface{} `json:"children"`
	} `json:"mapping"`
	ModerationResults []interface{} `json:"moderation_results"`
	CurrentNode       string        `json:"current_node"`
	ConversationId    string        `json:"conversation_id"`
}

func (api *GetConversation) Send() {
	request, _ := http.NewRequest(api.Method, api.URL+api.Rq.ConversationId, strings.NewReader(""))
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &api.Rs)
}

func (api *GetConversation) Default() {
	api.Name = "GetConversationWithId"
	api.Method = "GET"
	api.URL = BaseUrl + "/api/conversation/"
}
