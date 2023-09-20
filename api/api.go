package api

import (
	"net/http"
	"time"
)

const BaseUrl = "https://cg.zpaul.org"

type Api struct {
	Name   string
	Method string
	URL    string
}
type GPT interface {
	Send()
	Default()
}

func init() {
	http.DefaultClient.Timeout = 30 * time.Second
}
