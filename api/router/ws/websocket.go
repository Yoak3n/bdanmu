package ws

import (
	"bdanmu/app"
	"bdanmu/consts"
	"bdanmu/package/logger"
	"bdanmu/package/model"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  4094,
		WriteBufferSize: 4096,
		// websocket 真正的跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
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
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	if conn != nil {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("success"))
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
		for i, c := range serverHub.Clients {
			if c == client {
				serverHub.Clients = append(serverHub.Clients[:i], serverHub.Clients[i+1:]...)
				break
			}
		}
	}(conn)
	serverHub.Clients = append(serverHub.Clients, client)
	if conn != nil {
		_ = conn.WriteMessage(websocket.TextMessage, []byte("success"))
		logger.Logger.Infoln("2 client connected, total: ", len(serverHub.Clients)+1)
		go client.SendMessage()
		go client.ReadMessage()
		//go client.CloseClient()

		signal := <-client.CloseSignal
		logger.Logger.Infoln("client closed, signal: ")
		if signal {
			err := conn.Close()
			if err != nil {
				return
			}
		}
	} else {
		logger.Logger.Infoln("2 client closed, total: ", len(serverHub.Clients))
	}

}

func WriteMessage(message *model.Message) {
	data, err := json.Marshal(message)
	if err != nil {
		return
	}
	if clients := serverHub.Clients; len(clients) > 0 {
		for _, client := range clients {
			if client.Connection != nil {
				client.Message <- data
			}
		}
	}
}

func UpdateUser(user *model.User) {
	if user != nil {
		msg := &model.Message{
			Type: consts.USER_INFO,
			Data: user,
		}
		WriteMessage(msg)
	}
	if ctx := app.GetApp(); ctx != nil {
		runtime.EventsEmit(ctx.Ctx, "user", user)
	}
}
