package ws

import (
	"bdanmu/package/logger"
	"sync"

	"github.com/gorilla/websocket"
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
	for msg := range c.Message {
		c.Connection.WriteMessage(websocket.TextMessage, msg)
	}

}

func (c *Client) CloseClient() {

	signal := <-c.CloseSignal
	if signal {
		// maybe invalid address
		err := c.Connection.Close()
		if err != nil {
			return
		}
	}

}

func (c *Client) ReadMessage() {
	for {
		t, p, err := c.Connection.ReadMessage()
		if err != nil || t == -1 {
			c.CloseSignal <- true
			return
		}
		logger.Logger.Infoln(string(p))

	}
}
