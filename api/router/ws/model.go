package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Server struct {
	Clients []*Client
}

type Client struct {
	Connection  *websocket.Conn
	Mux         sync.Mutex
	Message     chan []byte
	CloseSignal chan bool
}

func (c *Client) SendMessage() {
	for {
		select {
		case msg := <-c.Message:
			c.Mux.Lock()
			c.Connection.WriteMessage(websocket.TextMessage, msg)
			c.Mux.Unlock()
		}
	}

}

func (c *Client) CloseClient() {
	for {
		select {
		case signal := <-c.CloseSignal:
			if signal {
				// maybe invalid address
				err := c.Connection.Close()
				if err != nil {
					return
				}
			}
		}
	}
}

func (c *Client) ReadMessage() {
	for {
		t, _, err := c.Connection.ReadMessage()
		if err != nil {
			c.Connection.Close()
			return
		}
		if t == -1 {
			c.Connection.Close()
			return
		}
		if t == 3000 {
			c.CloseSignal <- true
		}

	}
}
