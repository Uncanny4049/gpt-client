package client

import (
	"fmt"
	"github.com/Uncanny4049/gpt-client/api"
	"strings"
)

func SendQuestion(question string, title string) string {
	talk := api.PostConversationTalk{}
	talk.Default()
	talk.Rq = api.DefaultPostConversationTalkRQ(question)
	flag := true

	info := FindConversationWithTitle(title, 1, 20)
	if info.ConversationId != "" {
		// 仅为 talk.Rq 的部分属性赋值
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
	return talk.Rs.Message.Content.Parts[0]
}
func FindConversationWithTitle(title string, offset, limit int) api.GetConversationRS {
	conversations := api.GetConversations{
		Rq: api.GetConversationsRQ{
			Offset: offset,
			Limit:  limit,
		},
	}
	conversations.Default()
	conversations.Send()
	for _, item := range conversations.Rs.Items {
		if strings.EqualFold(item.Title, title) {
			conversation := api.GetConversation{Rq: api.GetConversationRQ{ConversationId: item.Id}}
			conversation.Default()
			conversation.Send()
			return conversation.Rs
		}
	}
	if conversations.Rs.Total > conversations.Rs.Limit*(conversations.Rs.Offset+1) {
		return FindConversationWithTitle(title, 1, conversations.Rs.Total)
	}
	return api.GetConversationRS{}
}
