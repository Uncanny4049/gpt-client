package client

import (
	"fmt"
	"github.com/Uncanny4049/gpt-client/api"
	"testing"
)

func TestSendQuestion(t *testing.T) {
	question := "DEMO"
	title := "2078788727"
	talk := api.PostConversationTalk{}
	talk.Default()
	talk.Rq = api.DefaultPostConversationTalkRQ(question)
	flag := true

	info := FindConversationWithTitle(title, 1, 20)
	if info.ConversationId != "" {
		talk.Rq.Model = info.Mapping[info.CurrentNode].Message.Metadata.ModelSlug
		talk.Rq.ParentMessageId = info.Mapping[info.CurrentNode].Message.Id
		talk.Rq.ConversationId = info.ConversationId
		flag = false
	}

	talk.Send()
	if flag {
		rename := api.PatConversation{
			Rq: api.PatConversationRQ{
				Title:          title,
				ConversationId: talk.Rs.ConversationId,
			},
		}
		rename.Default()
		rename.Send()
		if !rename.Rs.Success {
			fmt.Println("重命名失敗")
		}
	}
	t.Log(talk.Rs.Message.Content.Parts[0])
}
