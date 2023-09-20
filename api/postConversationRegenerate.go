package api

// PostConversationRegenerate /api/conversation/regenerate
// 让ChatGPT重新生成回复
type PostConversationRegenerate struct {
	Api
	Rq PostConversationRegenerateRQ
	Rs PostConversationRegenerateRS
}

// PostConversationRegenerateRQ POST /api/conversation/regenerate
type PostConversationRegenerateRQ struct {
}

// PostConversationRegenerateRS POST /api/conversation/regenerate
type PostConversationRegenerateRS struct {
}
