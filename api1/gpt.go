package api1

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var client http.Client = http.Client{
	Timeout: time.Second * 60,
}

func GetModels() GetModelsRS {
	resp, err := client.Get(Models)
	r := GetModelsRS{}
	if err != nil {
		return r
	}
	defer func(Body io.ReadCloser) {
		Body.Close()
	}(resp.Body)
	rs, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(rs, &r)
	return r
}

func Talk(rq TalkRQ) (TalkRQ, TalkContentRS) {
	jsonData, _ := json.Marshal(rq)
	response, _ := client.Post(ConversationTalk, "application/json", bytes.NewBuffer(jsonData))
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	rs := TalkRS{}
	_ = json.Unmarshal(body, &rs)
	return TalkRQ{
		Model:           rq.Model,
		MessageId:       uuid.NewString(),
		ParentMessageId: rs.Message.Id,
		ConversationId:  rs.ConversationId,
		Stream:          false,
	}, rs.Message.Content
}

// SetConversationTitle 修改名称
func SetConversationTitle(ConversationId string, data ConversationTitleRQ) {
	marshal, _ := json.Marshal(data)
	rq, _ := http.NewRequest("PATCH", Conversation+ConversationId, bytes.NewBuffer(marshal))
	rq.Header.Set("Content-Type", "application/json")
	rs, _ := client.Do(rq)
	rs.Body.Close()
}

// GetConversations 获取指定分页的对话
func GetConversations(qq string, offset, limit int) (conversationId string) {
	values := url.Values{
		"offset": []string{strconv.Itoa(offset)},
		"limit":  []string{strconv.Itoa(limit)},
	}
	resp, _ := client.Get(Conversations + "?" + values.Encode())
	defer resp.Body.Close()
	rs := GetConversationsRS{}
	rsStr, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(rsStr, &rs)
	for _, item := range rs.Items {
		if item.Title == qq {
			return item.Id
		}
	}
	if len(rs.Items) != limit {
		return ""
	} else {
		return GetConversations(qq, offset+1, limit)
	}

}

func GetOldConversationByTalk(conversationId string) TalkRQ {
	resp, _ := client.Get(Conversation + conversationId)
	defer resp.Body.Close()
	rs, _ := io.ReadAll(resp.Body)
	var info GetConversationRS
	_ = json.Unmarshal(rs, &info)
	return TalkRQ{
		Model:           info.Mapping[info.CurrentNode].Message.Metadata.ModelSlug,
		MessageId:       uuid.NewString(),
		ParentMessageId: info.Mapping[info.CurrentNode].Message.Id,
		ConversationId:  info.ConversationId,
		Stream:          false,
	}
}
