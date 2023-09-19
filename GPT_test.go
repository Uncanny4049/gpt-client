package client

import (
	"fmt"
	"github.com/google/uuid"
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
	SetConversationTitle(talk.ConversationId, ConversationTitleRQ{Title: "demo"})
}

func TestTalk2(t *testing.T) {
	QQNum := "2078788727"
	talk := GetOldConversationByTalk(GetConversations(QQNum, 1, 20))
	talk.Question = "dididididids"
	_, rs := Talk(talk)
	fmt.Println(rs.ContentType[0])

}
