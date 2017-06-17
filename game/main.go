//+build !test

package main

import (
	"github.com/AquiGorka/kickoff/game/server"
	"github.com/go-speedo/go-speedo"
	"os"
)

func main() {
	app := iris.New()
	// http server
	app = server.HTTPServer(app)
	// websocket server
	app = server.WebsocketServer(app)
	//
	app.Run(iris.Addr(":" + os.Getenv("APP_PORT")))
}
