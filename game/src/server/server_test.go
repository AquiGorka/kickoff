package server_test

import (
  "testing"
  "os"
  "gopkg.in/kataras/iris.v6"
  "game/server"
  "golang.org/x/net/websocket"
  "net/http"
)

var app *iris.Framework

func init() {
  app = iris.New()
  app = server.HttpServer(app)
  app = server.WebsocketServer(app)
  go app.Listen(":" + os.Getenv("APP_PORT"))
}

func TestSocketServer(t *testing.T) {
  _, err := websocket.Dial("ws://localhost:" + os.Getenv("APP_PORT")  + "/ws", "", "http://localhost/")
  if (err != nil) {
    t.Error("Error connecting to socket server: ", err)
  }
}

func TestHttpServer(t *testing.T) {
  _, err := http.Get("http://localhost:" + os.Getenv("APP_PORT"))
  if (err != nil) {
    t.Error("Error connecting to http server: ", err)
  }
}
