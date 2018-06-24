package main

import (
	"github.com/wsmdr/goPractice/wswebsocket"
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"bytes"
)

func main() {
	pushUrl := "http://127.0.0.1:12345/push"
	contentType := "application/json"

	for {
		pm := wswebsocket.PushMessage{
			UserID: "jack",
			Event: "topic1",
			Message: fmt.Sprintf("Hello in %s", time.Now().Format("2006-01-02 15:04:05.000")),
		}
		b, _ := json.Marshal(pm)

		http.DefaultClient.Post(pushUrl, contentType, bytes.NewReader(b))

		time.Sleep(10 * time.Second)
	}
}
