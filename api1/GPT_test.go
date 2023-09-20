package api1

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"testing"
)

func TestGetModels(t *testing.T) {
	models := GetModels()
	fmt.Println(models)
}

func TestTalk(t *testing.T) {
	talk, _ := Talk(TalkRQ{
		Question:        "使用Golang写一个HTTP的Post请求demo",
		Model:           "text-davinci-002-render-sha",
		MessageId:       uuid.NewString(),
		ParentMessageId: uuid.NewString(),
		Stream:          false,
	})
	SetConversationTitle(talk.ConversationId, ConversationTitleRQ{Title: "2078788727"})
}

func TestTalk2(t *testing.T) {
	QQNum := "2078788727"
	talk := GetOldConversationByTalk(GetConversations(QQNum, 1, 20))
	talk.Question = "dididididids"
	_, rs := Talk(talk)
	fmt.Println(rs.ContentType[0])
}

func Test(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		question string
	}{
		{name: "1", s: "1", question: "請幫我使用NodeJS寫一個HTTP請求範例"},
		{name: "2", s: "2", question: "請幫我使用Go寫一個HTTP請求範例"},
		{name: "3", s: "3", question: "請幫我使用Java寫一個HTTP請求範例"},
		{name: "4", s: "4", question: "請幫我使用Go寫一個HTTP請求範例"},
	}

	var wg sync.WaitGroup

	for _, test := range tests {
		wg.Add(1)
		tCopy := test
		go t.Run(test.name, func(t *testing.T) {
			defer wg.Done()

			ConversationId := GetConversations(tCopy.s, 1, 20)
			t.Log(ConversationId)
			var rq TalkRQ
			if ConversationId == "" {
				rq = TalkRQ{
					Model:           "text-davinci-002-render-sha",
					MessageId:       uuid.NewString(),
					ParentMessageId: uuid.NewString(),
					Stream:          false,
				}
			} else {
				rq = GetOldConversationByTalk(ConversationId)
			}
			rq.Question = tCopy.question
			t.Log(rq)
			rq, rs := Talk(rq)
			t.Log(rs)
			SetConversationTitle(rq.ConversationId, ConversationTitleRQ{Title: tCopy.s})
		})
	}
	wg.Wait()
}
