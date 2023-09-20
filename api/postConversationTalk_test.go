package api

import (
	"encoding/json"
	"testing"
)

func TestTalk(t *testing.T) {
	talk := PostConversationTalk{
		Rq: DefaultPostConversationTalkRQ("Go 如何撰寫測試方法"),
	}
	talk.Default()
	talk.Send()
	rs, _ := json.Marshal(talk.Rs)
	t.Log(string(rs))
	/* {"conversation_id":"43833584-0175-4ee9-b099-bcd34c8666fb","error":null,"message":{"author":{"metadata":{},"name":null,"role":"assistant"},"create_time":1695206805.968962,"end_turn":true,"id":"1b9144e4-0b6d-4e74-ab2e-966dc4b4c646","metadata":{"finish_details":{"stop_tokens":[100260],"type":"stop"},"is_complete":true,"message_type":"next","model_slug":"text-davinci-002-render-sha","parent_id":"6e824d4d-1d69-44dc-90c8-2fee2bfd4692"},"recipient":"all","status":"finished_successfully","update_time":null,"weight":1}}*/
}
