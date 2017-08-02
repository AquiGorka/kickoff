package server

import (
	"github.com/gorilla/websocket"
	"log"
)

func pingMsg(mt int, conn *websocket.Conn) {
	conn.WriteMessage(mt, []byte("pong"))
}

func onConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			break
		}
		// messages
		//log.Println("Message received: ", string(msg))
		if string(msg) == "ping" {
			pingMsg(mt, conn)
		}
	}
}
