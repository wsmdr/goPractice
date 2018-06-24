package wswebsocket

import (
	"github.com/gorilla/websocket"
	"io"
	"sync"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
)

type Conn struct {
	Conn *websocket.Conn
	AfterReadFunc func(messageType int, r io.Reader)
	BeforeCloseFunc func()

	once sync.Once
	id string
	stopCh chan struct{}
}

func (c *Conn) Write(p []byte) (n int, err error)  {
	select {
	case <-c.stopCh:
		return 0, errors.New("连接已关闭,不能写入")
	default:
		err := c.Conn.WriteMessage(websocket.TextMessage, p)
		if err != nil {
			return 0, err
		}
		return len(p), nil
	}
}

func (c *Conn) GetID() string {
	c.once.Do(func() {
		u := uuid.New()
		c.id = u.String()
	})

	return c.id
}

func (c *Conn) Listen() {
	c.Conn.SetCloseHandler(func(code int, text string) error {
		if c.BeforeCloseFunc != nil {
			c.BeforeCloseFunc()
		}

		if err := c.Close(); err != nil {
			log.Println(err)
		}
		message :=websocket.FormatCloseMessage(code, "")
		c.Conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})
ReadLoop:
	for  {
		select {
		case <-c.stopCh:
			break ReadLoop
		default:
			messageType, r, err := c.Conn.NextReader()
			if err != nil {
				break ReadLoop
			}
			if c.AfterReadFunc != nil {
				c.AfterReadFunc(messageType, r)
			}
		}

	}
}

func (c *Conn) Close() error {
	select {
	case <-c.stopCh:
		return errors.New("连接已经被关闭")
	default:
		c.Conn.Close()
		close(c.stopCh)
		return nil
	}
}

func NewConn(conn *websocket.Conn) *Conn {
	return &Conn{
		Conn: conn,
		stopCh: make(chan struct{}),
	}
}