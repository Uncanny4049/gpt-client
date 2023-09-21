package api

import (
	"net/http"
	"time"
)

const BaseUrl = "http://127.0.0.1:8001"

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
