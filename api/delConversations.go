package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// DelConversations /api/conversations
// 删除所有会话
type DelConversations struct {
	Api
	Rq DelConversationsRQ
	Rs DelConversationsRS
}

// DelConversationsRQ Delete /api/conversations
type DelConversationsRQ struct {
}

// DelConversationsRS Delete /api/conversations
type DelConversationsRS struct {
	Success bool `json:"success"`
}

func (api *DelConversations) Send() {
	request, _ := http.NewRequest(api.Method, api.URL, strings.NewReader(""))
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &api.Rs)

}

func (api *DelConversations) Default() {
	api.Name = "ClearAllConversations"
	api.Method = "DELETE"
	api.URL = BaseUrl + "/api/conversations"
}
