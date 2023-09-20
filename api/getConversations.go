package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GetConversations Get /api/conversations
// 以分页方式列出会话列表
type GetConversations struct {
	Api
	Rq GetConversationsRQ
	Rs GetConversationsRS
}

// GetConversationsRQ Get /api/conversations
type GetConversationsRQ struct {
	Offset int // 頁數
	Limit  int // 分頁顯示數量
}

// GetConversationsRS Get /api/conversations
type GetConversationsRS struct {
	Items                   []GetConversationInfo `json:"items"`
	Total                   int                   `json:"total"`
	Limit                   int                   `json:"limit"`
	Offset                  int                   `json:"offset"`
	HasMissingConversations bool                  `json:"has_missing_conversations"`
}
type GetConversationInfo struct {
	Id                     string      `json:"id"`
	Title                  string      `json:"title"`
	CreateTime             time.Time   `json:"create_time"`
	UpdateTime             time.Time   `json:"update_time"`
	Mapping                interface{} `json:"mapping"`
	CurrentNode            interface{} `json:"current_node"`
	ConversationTemplateId interface{} `json:"conversation_template_id"`
}

func (api *GetConversations) Send() {
	request, _ := http.NewRequest(api.Method, api.URL+"?"+url.Values{
		"offset": []string{strconv.Itoa(api.Rq.Offset)},
		"limit":  []string{strconv.Itoa(api.Rq.Limit)},
	}.Encode(), strings.NewReader(""))
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	_ = json.Unmarshal(body, &api.Rs)
}

func (api *GetConversations) Default() {
	api.Name = "GetConversations"
	api.Method = "GET"
	api.URL = BaseUrl + "/api/conversations"
}
