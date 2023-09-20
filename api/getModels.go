package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// GetModels Get /api/models
// 列出账号可用的模型
type GetModels struct {
	Api
	Rq GetModelsRQ
	Rs GetModelsRS
}

// GetModelsRQ Get /api/models
type GetModelsRQ struct {
}

// GetModelsRS Get /api/models
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

func (api *GetModels) Send() {
	request, _ := http.NewRequest(api.Method, api.URL, strings.NewReader(""))
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	_ = json.Unmarshal(body, &api.Rs)
}

func (api *GetModels) Default() {
	api.Name = "GetModels"
	api.Method = "GET"
	api.URL = BaseUrl + "/api/models"
}
