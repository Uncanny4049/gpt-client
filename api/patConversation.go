package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// PatConversation /api/conversation/<conversation_id>
// 通过会话ID设置指定的会话标题
type PatConversation struct {
	Api
	Rq PatConversationRQ
	Rs PatConversationRS
}

// PatConversationRQ Patch /api/conversation/<conversation_id>
type PatConversationRQ struct {
	Title          string `json:"title"` // 新标题
	ConversationId string `json:"-"`
}

// PatConversationRS Patch /api/conversation/<conversation_id>
type PatConversationRS struct {
	Detail  string
	Success bool
}

func (api *PatConversation) Send() {
	data, _ := json.Marshal(api.Rq)
	request, _ := http.NewRequest(api.Method, api.URL+api.Rq.ConversationId, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &api.Rs)
}

func (api *PatConversation) Default() {
	api.Name = "SetConversationTitleWithId"
	api.Method = "PATCH"
	api.URL = BaseUrl + "/api/conversation/"
}
