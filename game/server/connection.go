package server

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type msg struct {
	t byte
}

func pingMsg(conn *websocket.Conn) {
	if err := conn.WriteJSON("pong"); err != nil {
		fmt.Println(err)
	}
}

func onConnection(conn *websocket.Conn) {
	for {
		m := msg{}
		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}
		fmt.Printf("Got message: %#v\n", m)
		pingMsg(conn)
	}
}
