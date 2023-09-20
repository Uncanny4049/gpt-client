package api

// PostConversationGoon /api/conversation/goon
// 让ChatGPT讲之前的恢复继续下去
type PostConversationGoon struct {
	Api
	Rq PostConversationGoonRQ
	Rs PostConversationGoonRS
}

// PostConversationGoonRQ POST /api/conversation/goon
type PostConversationGoonRQ struct {
}

// PostConversationGoonRS POST /api/conversation/goon
type PostConversationGoonRS struct {
}
