package api

import (
	"encoding/json"
	"testing"
)

func TestGetConversations(t *testing.T) {
	conversations := GetConversations{
		Rq: GetConversationsRQ{
			Offset: 1,
			Limit:  2,
		},
	}
	conversations.Default()
	conversations.Send()
	rs, _ := json.Marshal(conversations.Rs)
	t.Log(string(rs))
	/**{"items":[{"id":"a2f677ab-f1fa-4537-900b-81f594de3ac8","title":"New chat","create_time":"2023-09-20T10:44:38.782672Z","update_time":"2023-09-20T10:44:38Z","mapping":null,"current_node":null,"conversation_template_id":null},{"id":"bca98c01-17f7-448c-89eb-81f55ab8ee7d","title":"4","create_time":"2023-09-20T09:31:01.440924Z","update_time":"2023-09-20T09:31:15Z","mapping":null,"current_node":null,"conversation_template_id":null},{"id":"99e18132-dd1f-4d0f-9587-5b3e5f94b364","title":"1","create_time":"2023-09-20T09:31:01.41352Z","update_time":"2023-09-20T09:31:17Z","mapping":null,"current_node":null,"conversation_template_id":null},{"id":"5f71e6a5-e57e-40f3-b0bd-fa8dd82bd1f6","title":"3","create_time":"2023-09-20T09:31:01.366103Z","update_time":"2023-09-20T09:31:17Z","mapping":null,"current_node":null,"conversation_template_id":null},{"id":"8db4f7d2-8099-4b9e-aaa0-3a049335276d","title":"2","create_time":"2023-09-20T09:31:01.351154Z","update_time":"2023-09-20T09:31:16Z","mapping":null,"current_node":null,"conversation_template_id":null}],"total":6,"limit":20,"offset":1,"has_missing_conversations":false}*/
}
