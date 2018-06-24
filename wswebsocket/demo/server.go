package main

import (
	"github.com/wsmdr/goPractice/wswebsocket"
	"net/http"
)

func main() {
	server := wswebsocket.NewServer(":12345")

	server.WSPATH = "/ws"
	server.PushPath = "/push"

	server.AuthToken = func(token string) (userId string, ok bool) {
		if token == "aaa" {
			return "jack", true
		}

		return "", false
	}

	server.PushAuth = func(r *http.Request) bool {

		return true
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}