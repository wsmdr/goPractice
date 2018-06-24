package wswebsocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"io"
	"encoding/json"
	"log"
	"github.com/pkg/errors"
	"fmt"
	"strings"
)

type websocketHandler struct {
	upgrader *websocket.Upgrader
	binder *binder

	callUserIDFunc func(token string) (userId string, ok bool)
}

type RegisterMessage struct {
	Token string
	Event string
}

func (wh *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	wsConn, err := wh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer wsConn.Close()

	conn := NewConn(wsConn)
	conn.AfterReadFunc = func(messageType int, r io.Reader) {
		var rm RegisterMessage
		decoder := json.NewDecoder(r)
		if err := decoder.Decode(&rm); err != nil {
			return
		}
		userID := rm.Token
		if wh.callUserIDFunc != nil {
			uID, ok := wh.callUserIDFunc(rm.Token)
			if !ok {
				return
			}
			userID = uID
		}

		wh.binder.Bind(userID, rm.Event, conn)
	}
	conn.BeforeCloseFunc = func() {
		wh.binder.Unbind(conn)
	}
	conn.Listen()
}

func (wh *websocketHandler) closeConns(userID, event string) (int, error) {
	conns, err := wh.binder.FilterConn(userID, event)
	if err != nil {
		return 0,err
	}

	cnt := 0
	for i := range conns {
		if err := wh.binder.Unbind(conns[i]); err != nil {
			log.Printf("conn unbind fail: %v", err)
			continue
		}

		if err := conns[i].Close(); err != nil {
			log.Printf("conn close fail: %v", err)
			continue
		}
		cnt++
	}
	return cnt, nil
}

var ErrRequestIllegal = errors.New("request data illegal")

type pushHandler struct {
	authFunc func(r *http.Request) bool
	binder *binder
}

func (s *pushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if s.authFunc != nil {
		if ok := s.authFunc(r); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	var pm PushMessage
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrRequestIllegal.Error()))
		return
	}

	if pm.UserID == "" || pm.Event == "" || pm.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrRequestIllegal.Error()))
		return
	}

	cnt, err := s.push(pm.UserID, pm.Event, pm.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result := strings.NewReader(fmt.Sprintf("message send to %d clients", cnt))
	io.Copy(w, result)
}

func (s *pushHandler) push(userID, event, message string) (int, error) {
	if userID == "" || event == "" || message == "" {
		return 0, errors.New("parameters(userID, event, message) can't be empty")
	}

	conns, err := s.binder.FilterConn(userID, event)
	if err != nil {
		return 0, fmt.Errorf("filter conn fail: %v", err)
	}
	cnt := 0
	for i := range conns {
		_, err := conns[i].Write([]byte(message))
		if err != nil {
			s.binder.Unbind(conns[i])
			continue
		}
		cnt++
	}

	return cnt, nil
}

type PushMessage struct {
	UserID string `json:"userId"`
	Event string
	Message string
}
