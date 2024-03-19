package ws

import (
	"bdanmu/package/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	serverHub *Server
)

func init() {
	serverHub = &Server{
		Clients: make([]*Client, 0),
	}
}

func GetHub() *Server {
	return serverHub
}

func RegisterClient(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	origin := c.Request.Header.Get("Origin")
	if upgrader.CheckOrigin == nil || upgrader.CheckOrigin(c.Request) {
		if origin != "" {
			logger.Logger.Println("origin:", origin)
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Origin", origin)
		}
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		client := &Client{
			Connection:  conn,
			Message:     make(chan []byte),
			CloseSignal: make(chan bool),
		}
		defer func(conn *websocket.Conn) {
			if conn != nil {
				_ = conn.Close()
			}
		}(conn)
		serverHub.Clients = append(serverHub.Clients, client)
		if conn == nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("success"))
		}
		go client.SendMessage()
		go client.ReadMessage()
		//go client.CloseClient()
		select {
		case signal := <-client.CloseSignal:
			if signal {
				// maybe invalid address
				err := conn.Close()
				if err != nil {
					return
				}
			}
		}
	}
}

func WriteMessage(data []byte) {
	if clients := serverHub.Clients; len(clients) > 0 {
		for _, client := range clients {
			client.Message <- data
		}
	}
}
