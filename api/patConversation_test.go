package api

import (
	"encoding/json"
	"testing"
)

func TestPatConversation(t *testing.T) {
	conversation := PatConversation{}
	conversation.Default()
	conversation.Rq.ConversationId = "99e18132-dd1f-4d0f-9187-5b3e5f94b364"
	conversation.Rq.Title = "demo"
	conversation.Send()
	rs, _ := json.Marshal(conversation.Rs)
	t.Log(string(rs))
	/**{"Detail":""}*/
	/**{"Detail":"Invalid conversation 43833584-0175-4ee9-b099-bcd34c866fb"}*/
}
