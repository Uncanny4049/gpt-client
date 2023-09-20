package api

// DelConversation /api/conversation/<conversation_id>
// 通过会话ID删除指定会话
type DelConversation struct {
	Api
	Rq DelConversationRQ
	Rs DelConversationRS
}

// DelConversationRQ DELETE /api/conversation/<conversation_id>
type DelConversationRQ struct {
	ConversationId string
}

// DelConversationRS DELETE /api/conversation/<conversation_id>
type DelConversationRS struct {
}
