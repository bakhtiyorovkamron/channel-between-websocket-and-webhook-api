package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

// type Data struct {
// 	IsNew   bool
// 	Message interface{}
// }

// var data *Data

func (h *handlerV1) GetLocation(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("err :", err)
		return
	}
	defer conn.Close()
	for {
		resp := <-myChannel
		conn.WriteJSON(resp)
	}
}

// Path: api/handler/socket.go
