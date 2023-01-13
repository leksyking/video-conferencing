package handlers

import (
	"fmt"
	"os"

	w "video-conferencing/pkg/webrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func RoomWebsocket(c *websocket.Conn) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return nil
	}
	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
	//room connection
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}
	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
	return c.Render("peer", fiber.Map{
		"RommWebsocketAddress":
		"RoomLink":
		
	})
}

func createOrGetRoom(uuid string) (string, string, *w.Room) {

}

func RoomViewerWebsocket(c *websocket.Conn) {

}
func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}

type websocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
