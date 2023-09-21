package main

import (
	"encoding/json"
	"fmt"
	"github.com/Uncanny4049/gpt-client/client"
	"net/http"
)

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1")
		r.ParseForm()
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(r.Form.Get("question"))
		fmt.Println(r.Form.Get("title"))
		responseData, _ := json.Marshal(ResponseData{
			Code:    200,
			Message: client.SendQuestion(r.Form.Get("question"), r.Form.Get("title")),
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseData)
		return
	})
	address := "127.0.0.1:8081"
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Printf("HTTP server failed to start: %v", err)
	}
}
