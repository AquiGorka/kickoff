package main

import(
  "os"
  "game/server"
)

func main() {
  // run http server
  server.HttpServer(os.Getenv("APP_PORT"))
}
