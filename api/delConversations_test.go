package api

import (
	"encoding/json"
	"testing"
)

func TestDelConversations(t *testing.T) {
	conversation := DelConversations{}
	conversation.Default()
	conversation.Send()
	rs, _ := json.Marshal(conversation.Rs)
	t.Log(string(rs))
}
