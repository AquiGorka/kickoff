package server

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

// HTTPServer sets the handlers for the http routes
func HTTPServer(router *mux.Router) *mux.Router {
	// ping
	router.HandleFunc("/ping", pingHandler)
	//
	return router
}

// WebsocketServer sets the handler for websocket connections
func WebsocketServer(router *mux.Router) *mux.Router {
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}
		go onConnection(conn)
	})
	//
	return router
}
