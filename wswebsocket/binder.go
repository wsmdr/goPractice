package wswebsocket

import (
	"sync"
	"errors"
	"fmt"
)

type eventConn struct {
	Event string
	Conn *Conn
}

type binder struct {
	mu sync.RWMutex
	userID2EventConnMap map[string]*[]eventConn

	connID2UserIDMap map[string]string
}

func (b *binder) Bind(userID, event string, conn *Conn) error {
	if userID == "" {
		return errors.New("userID 不能为空")
	}
	if event == "" {
		return errors.New("event 不能为空")
	}
	if conn == nil {
		return errors.New("conn can't be nil")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if eConns, ok := b.userID2EventConnMap[userID]; ok {
		for i := range *eConns {
			if (*eConns)[i].Conn == conn {
				return nil
			}
		}
		newEConns := append(*eConns, eventConn{event, conn})
		b.userID2EventConnMap[userID] = &newEConns
	} else {
		b.userID2EventConnMap[userID] = &[]eventConn{{event, conn}}
	}
	b.connID2UserIDMap[conn.GetID()] = userID
	return nil
}

func (b *binder) Unbind(conn *Conn) error {
	if conn == nil {
		return errors.New("conn can't be nil")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	userId, ok := b.connID2UserIDMap[conn.GetID()]
	if !ok {
		return fmt.Errorf("can't find userId by connID: %s", conn.GetID())
	}

	if eConns, ok := b.userID2EventConnMap[userId]; ok {
		for i := range *eConns {
			if (*eConns)[i].Conn == conn {
				newEConns := append((*eConns)[:i], (*eConns)[i+1:]...)
				b.userID2EventConnMap[userId] = &newEConns
				delete(b.connID2UserIDMap, conn.GetID())

				if len(newEConns) == 0 {
					delete(b.userID2EventConnMap, userId)
				}
				return nil
			}
		}
		return fmt.Errorf("can't find the conn of ID: %s", conn.GetID())
	}
	return fmt.Errorf("can;t find the eventConns by userID: %s", userId)
}

func (b *binder) FindConn(connId string) (*Conn, bool) {
	if connId == "" {
		return nil, false
	}
	userID, ok := b.connID2UserIDMap[connId]

	if ok {
		if eConns, ok := b.userID2EventConnMap[userID]; ok {
			for i := range *eConns {
				if (*eConns)[i].Conn.GetID() == connId {
					return (*eConns)[i].Conn, true
				}
			}
		}

		return nil, false
	}

	for _, eConns := range b.userID2EventConnMap {
		for i := range *eConns {
			if (*eConns)[i].Conn.GetID() == connId {
				return (*eConns)[i].Conn, true
			}

		}

	}
	return nil, false
}

func (b *binder) FilterConn(userID, event string) ([]*Conn, error) {
	if userID == "" {
		return nil, errors.New("userId can't be empty")
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	if eConns, ok := b.userID2EventConnMap[userID]; ok {
		ecs := make([]*Conn, 0, len(*eConns))
		for i := range *eConns {
			if event == "" || (*eConns)[i].Event == event {
				ecs = append(ecs, (*eConns)[i].Conn)
			}
		}
		return ecs, nil
	}

	return []*Conn{}, nil
}
