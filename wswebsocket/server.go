package wswebsocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"strings"
	"errors"
)

const (
	serverDefaultWSPath = "/ws"
	serverDefaultPushPath = "/push"
)

var defaultUpgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	Addr string
	WSPATH string
	PushPath string
	Upgrader *websocket.Upgrader

	AuthToken func(token string) (userId string, ok bool)

	PushAuth func(r *http.Request) bool

	wh *websocketHandler
	ph *pushHandler
}

func (s *Server) ListenAndServe() error {
	b := &binder{
		userID2EventConnMap: make(map[string]*[]eventConn),
		connID2UserIDMap: make(map[string]string),
	}

	wh := websocketHandler{
		upgrader: defaultUpgrader,
		binder: b,
	}
	if s.Upgrader != nil {
		wh.upgrader = s.Upgrader
	}
	if s.AuthToken != nil {
		wh.callUserIDFunc = s.AuthToken
	}

	s.wh = &wh
	http.Handle(s.WSPATH, s.wh)

	ph := pushHandler{
		binder: b,
	}
	if s.PushAuth != nil {
		ph.authFunc = s.PushAuth
	}
	s.ph = &ph
	http.Handle(s.PushPath, s.ph)

	return http.ListenAndServe(s.Addr, nil)
}

func (s *Server) Push(userID, event, message string) (int, error) {
	return s.ph.push(userID, event, message)
}

func (s *Server) Drop(userID, event string) (int, error) {
	return s.wh.closeConns(userID, event)
}

func (s Server) check() error {
	if !checkPath(s.WSPATH) {
		return fmt.Errorf("WSPath: %s not illegal", s.WSPATH)
	}
	if !checkPath(s.PushPath) {
		return fmt.Errorf("PushPath: %s not illegal", s.PushPath)
	}
	if s.WSPATH == s.PushPath {
		return errors.New("WSPath is equal to PushPath")
	}

	return nil
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
		WSPATH: serverDefaultWSPath,
		PushPath: serverDefaultPushPath,
	}
}

func checkPath(path string) bool {
	if path != "" && strings.HasPrefix(path, "/") {
		return false
	}
	return true
}

