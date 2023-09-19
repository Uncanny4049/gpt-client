package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	resp, err := client.Get("http://cg.zpaul.org/api/models")
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
	fmt.Println(string(jsonData))
	resp, _ := client.Post(ConversationTalk, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()
	rs, _ := io.ReadAll(resp.Body)
	fmt.Println(string(rs))
	talkrs := TalkRS{}
	_ = json.Unmarshal(rs, &talkrs)

	return TalkRQ{
		Model:           rq.Model,
		MessageId:       uuid.NewString(),
		ParentMessageId: talkrs.Message.Id,
		ConversationId:  talkrs.ConversationId,
		Stream:          false,
	}, talkrs.Message.Content
}

// SetConversationTitle 修改名称
func SetConversationTitle(ConversationId string, data ConversationTitleRQ) {
	marshal, _ := json.Marshal(data)
	rq, _ := http.NewRequest("PATCH", Conversation+ConversationId, bytes.NewBuffer(marshal))
	rq.Header.Set("Content-Type", "application/json")
	fmt.Println(rq)
	rs, _ := client.Do(rq)
	defer rs.Body.Close()
	s, _ := io.ReadAll(rs.Body)
	fmt.Println(string(s))
}

// GetConversations 获取指定分页的对话
func GetConversations(qq string, offset, limit int) (conversationId string) {
	values := url.Values{
		"offset": []string{strconv.Itoa(offset)},
		"limit":  []string{strconv.Itoa(limit)},
	}
	resp, _ := client.Get(Conversations + "?" + values.Encode())
	fmt.Println(resp.Request.URL)
	defer resp.Body.Close()
	rs := GetConversationsRS{}
	rsStr, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(rsStr, &rs)
	fmt.Println(rs)
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
	var info GetConversationRQ
	_ = json.Unmarshal(rs, &info)
	return TalkRQ{
		Model:           info.Mapping[info.CurrentNode].Message.Metadata.ModelSlug,
		MessageId:       uuid.NewString(),
		ParentMessageId: info.Mapping[info.CurrentNode].Message.Id,
		ConversationId:  info.ConversationId,
		Stream:          false,
	}
}
