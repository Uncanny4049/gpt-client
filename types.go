package client

import "time"

type GetModelsRS struct {
	Models []struct {
		Slug         string   `json:"slug"`
		MaxTokens    int      `json:"max_tokens"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Tags         []string `json:"tags"`
		Capabilities struct {
		} `json:"capabilities"`
		ProductFeatures struct {
		} `json:"product_features"`
	} `json:"models"`
	Categories []struct {
		Category             string `json:"category"`
		HumanCategoryName    string `json:"human_category_name"`
		SubscriptionLevel    string `json:"subscription_level"`
		DefaultModel         string `json:"default_model"`
		BrowsingModel        string `json:"browsing_model"`
		CodeInterpreterModel string `json:"code_interpreter_model"`
		PluginsModel         string `json:"plugins_model"`
	} `json:"categories"`
}

type TalkRQ struct {
	Question        string `json:"prompt"`
	Model           string `json:"model"`
	MessageId       string `json:"message_id"`
	ParentMessageId string `json:"parent_message_id"`
	ConversationId  string `json:"conversation_id,omitempty"`
	Stream          bool   `json:"stream" default:"false"`
}

type TalkRS struct {
	ConversationId string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
	Message        struct {
		Author struct {
			Metadata struct {
			} `json:"metadata"`
			Name interface{} `json:"name"`
			Role string      `json:"role"`
		} `json:"author"`
		Content    TalkContentRS `json:"content"`
		CreateTime float64       `json:"create_time"`
		EndTurn    bool          `json:"end_turn"`
		Id         string        `json:"id"`
		Metadata   struct {
			FinishDetails struct {
				StopTokens []int  `json:"stop_tokens"`
				Type       string `json:"type"`
			} `json:"finish_details"`
			IsComplete  bool   `json:"is_complete"`
			MessageType string `json:"message_type"`
			ModelSlug   string `json:"model_slug"`
			ParentId    string `json:"parent_id"`
		} `json:"metadata"`
		Recipient  string      `json:"recipient"`
		Status     string      `json:"status"`
		UpdateTime interface{} `json:"update_time"`
		Weight     float64     `json:"weight"`
	} `json:"message"`
}

type TalkContentRS struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

type ConversationTitleRQ struct {
	Title string `json:"title"`
}
type GetConversationRQ struct {
	Title             string                          `json:"title"`
	CreateTime        float64                         `json:"create_time"`
	UpdateTime        float64                         `json:"update_time"`
	Mapping           map[string]GetConversationRqMap `json:"mapping"`
	ModerationResults []interface{}                   `json:"moderation_results"`
	CurrentNode       string                          `json:"current_node"`
	ConversationId    string                          `json:"conversation_id"`
}

type GetConversationRqMap struct {
	Id      string `json:"id"`
	Message struct {
		Id     string `json:"id"`
		Author struct {
			Role     string `json:"role"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"author"`
		CreateTime float64 `json:"create_time"`
		Content    struct {
			ContentType string   `json:"content_type"`
			Parts       []string `json:"parts"`
		} `json:"content"`
		Status   string  `json:"status"`
		EndTurn  bool    `json:"end_turn"`
		Weight   float64 `json:"weight"`
		Metadata struct {
			FinishDetails struct {
				Type       string `json:"type"`
				StopTokens []int  `json:"stop_tokens"`
			} `json:"finish_details"`
			IsComplete bool   `json:"is_complete"`
			ModelSlug  string `json:"model_slug"`
			ParentId   string `json:"parent_id"`
			Timestamp  string `json:"timestamp_"`
		} `json:"metadata"`
		Recipient string `json:"recipient"`
	} `json:"message"`
	Parent   string        `json:"parent"`
	Children []interface{} `json:"children"`
}

type GetConversationsRS struct {
	Items []struct {
		Id                     string      `json:"id"`
		Title                  string      `json:"title"`
		CreateTime             time.Time   `json:"create_time"`
		UpdateTime             time.Time   `json:"update_time"`
		Mapping                interface{} `json:"mapping"`
		CurrentNode            interface{} `json:"current_node"`
		ConversationTemplateId interface{} `json:"conversation_template_id"`
	} `json:"items"`
	Total                   int  `json:"total"`
	Limit                   int  `json:"limit"`
	Offset                  int  `json:"offset"`
	HasMissingConversations bool `json:"has_missing_conversations"`
}
