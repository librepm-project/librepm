package http

import (
	"net/http"

	"libs/jwt_utils"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketControllerInterface interface {
	Connect(w http.ResponseWriter, r *http.Request)
}

type WebSocketController struct {
	Hub *Hub
}

func (c WebSocketController) Connect(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	tokenInfo := jwt_utils.GetTokenInfo(token)
	if tokenInfo == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		userID: tokenInfo.UserID,
		conn:   conn,
		send:   make(chan []byte, 64),
		hub:    c.Hub,
	}
	c.Hub.register(client)
	go client.writePump()
	client.readPump()
}
