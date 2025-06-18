package websocket

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)

// Client represents a connected WebSocket client
type Client struct {
	Conn   *websocket.Conn
	UserID string
	Send   chan []byte
	Hub    *Hub
}

// Hub maintains the set of active clients and broadcasts messages
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// NewHub creates a new hub instance
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(hub *Hub) websocket.Handler {
	return func(c *websocket.Conn) {
		// Create new client
		client := &Client{
			Conn:   c,
			UserID: c.Query("user_id"), // Get user ID from query parameter
			Send:   make(chan []byte, 256),
			Hub:    hub,
		}

		// Register client
		client.Hub.register <- client

		// Handle client disconnection
		defer func() {
			client.Hub.unregister <- client
			client.Conn.Close()
		}()

		// Start goroutine to read messages
		go func() {
			for {
				_, message, err := client.Conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Printf("error: %v", err)
					}
					break
				}

				// Broadcast message to all clients
				client.Hub.broadcast <- message
			}
		}()

		// Start goroutine to write messages
		for {
			select {
			case message, ok := <-client.Send:
				if !ok {
					// The hub closed the channel
					client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				w, err := client.Conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				w.Write(message)

				// Add queued messages to the current websocket message
				n := len(client.Send)
				for i := 0; i < n; i++ {
					w.Write([]byte{'\n'})
					w.Write(<-client.Send)
				}

				if err := w.Close(); err != nil {
					return
				}
			}
		}
	}
}

// Message represents a WebSocket message
type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// BroadcastMessage broadcasts a message to all connected clients
func (h *Hub) BroadcastMessage(msgType string, payload interface{}) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message := Message{
		Type:    msgType,
		Payload: payloadBytes,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	h.broadcast <- messageBytes
	return nil
}
