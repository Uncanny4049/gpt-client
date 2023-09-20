package api

// PostConversationGenTitle /api/conversation/gen_title/<conversation_id>
// 自动生成指定新会话的标题，通常首次问答后调用
type PostConversationGenTitle struct {
	Api
	Rq PostConversationGenTitleRQ
	Rs PostConversationGenTitleRS
}

// PostConversationGenTitleRQ POST /api/conversation/gen_title/<conversation_id>
type PostConversationGenTitleRQ struct {
	Model     string // 对话所使用的模型
	MessageId string // ChatGPT回复的那条消息的ID
}

// PostConversationGenTitleRS POST /api/conversation/gen_title/<conversation_id>
type PostConversationGenTitleRS struct {
}
