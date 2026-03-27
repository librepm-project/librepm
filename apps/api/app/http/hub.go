package http

import (
	"encoding/json"
	"sync"

	"apps/api/app/domain"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsMessage struct {
	Type    string      `json:"type"`
	Channel string      `json:"channel,omitempty"`
	Entity  string      `json:"entity,omitempty"`
	ID      string      `json:"id,omitempty"`
	Data    interface{} `json:"data"`
}

type inboundMessage struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}

type Client struct {
	userID uuid.UUID
	conn   *websocket.Conn
	send   chan []byte
	hub    *Hub
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}
}

func (c *Client) readPump() {
	defer c.hub.unregister(c)
	for {
		_, raw, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		var msg inboundMessage
		if json.Unmarshal(raw, &msg) == nil {
			c.hub.handleClientMessage(c, msg)
		}
	}
}

type Hub struct {
	mu       sync.RWMutex
	users    map[uuid.UUID][]*Client
	channels map[string][]*Client
}

func NewHub() *Hub {
	return &Hub{
		users:    make(map[uuid.UUID][]*Client),
		channels: make(map[string][]*Client),
	}
}

func (h *Hub) register(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.users[client.userID] = append(h.users[client.userID], client)
}

func (h *Hub) unregister(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	userClients := h.users[client.userID]
	for i, c := range userClients {
		if c == client {
			h.users[client.userID] = append(userClients[:i], userClients[i+1:]...)
			break
		}
	}
	for ch, chClients := range h.channels {
		for i, c := range chClients {
			if c == client {
				h.channels[ch] = append(chClients[:i], chClients[i+1:]...)
				break
			}
		}
	}
	close(client.send)
}

func (h *Hub) handleClientMessage(client *Client, msg inboundMessage) {
	switch msg.Type {
	case "subscribe":
		h.subscribeClient(msg.Channel, client)
	case "unsubscribe":
		h.unsubscribeClient(msg.Channel, client)
	}
}

func (h *Hub) subscribeClient(channel string, client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.channels[channel] = append(h.channels[channel], client)
}

func (h *Hub) unsubscribeClient(channel string, client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	clients := h.channels[channel]
	for i, c := range clients {
		if c == client {
			h.channels[channel] = append(clients[:i], clients[i+1:]...)
			return
		}
	}
}

func trySend(c *Client, msg []byte) {
	defer func() { recover() }() //nolint:errcheck
	select {
	case c.send <- msg:
	default:
	}
}

func (h *Hub) pushToUser(userID uuid.UUID, msg []byte) {
	h.mu.RLock()
	clients := make([]*Client, len(h.users[userID]))
	copy(clients, h.users[userID])
	h.mu.RUnlock()
	for _, c := range clients {
		trySend(c, msg)
	}
}

func (h *Hub) Broadcast(channel string, msg []byte) {
	h.mu.RLock()
	clients := make([]*Client, len(h.channels[channel]))
	copy(clients, h.channels[channel])
	h.mu.RUnlock()
	for _, c := range clients {
		trySend(c, msg)
	}
}

// PushToUser implements domain.NotificationPusher
func (h *Hub) PushToUser(userID uuid.UUID, n *domain.NotificationModel) {
	msg, err := json.Marshal(WsMessage{
		Type: "notification",
		Data: NotificationSerializer{}.ModelToResponse(*n),
	})
	if err != nil {
		return
	}
	h.pushToUser(userID, msg)
}
