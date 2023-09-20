package api

import (
	"encoding/json"
	"testing"
)

func TestGetConversation(t *testing.T) {
	conversation := GetConversation{Rq: GetConversationRQ{ConversationId: "91bcb491-1e1e-4386-b22b-c192f1dc8901"}}
	conversation.Default()
	conversation.Send()
	data, _ := json.Marshal(conversation.Rs)
	t.Log(string(data))

}
