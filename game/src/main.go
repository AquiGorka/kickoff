package main

import(
  "os"
  "game/server"
  "gopkg.in/kataras/iris.v6"
)

func main() {
  app := iris.New()
  // http server
  app = server.HttpServer(app)
  // websocket server
  app = server.WebsocketServer(app)
  //
  app.Listen(":" + os.Getenv("APP_PORT"))
}
