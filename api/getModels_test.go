package api

import (
	"encoding/json"
	"testing"
)

func TestGetModels(t *testing.T) {
	models := GetModels{}
	models.Default()
	models.Send()
	rs, _ := json.Marshal(models.Rs)
	t.Log(string(rs))

	/*{"models":[{"slug":"text-davinci-002-render-sha","max_tokens":8191,"title":"Default(GPT-3.5)","description":"Our fastest model, great for most everydaytasks.","tags":["gpt3.5"],"capabilities":{},"product_features":{}}],"categories":[{"category":"gpt_3.5","human_category_name":"GPT-3.5","subscription_level":"free","default_model":"text-davinci-002-render-sha","browsing_model":"text-davinci-002-render-sha-browsing","code_interpreter_model":"text-davinci-002-render-sha-code-interpreter","plugins_model":"text-davinci-002-render-sha-plugins"}]}*/
}
