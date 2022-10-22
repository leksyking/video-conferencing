package chat

import (
	"time"

	"github.com/fasthttp/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *Client) readPump() {

}
func (c *Client) writePump() {

}

func PeerChatConn(c *websocket.Conn, hub *Hub) {
	client := &Client{
		Hub:  hub,
		Conn: c,
		Send: make(chan []byte, 256),
	}
	client.Hub.register <- client
	go client.writePump()
	client.readPump()
}
