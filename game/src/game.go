package main

import (
  "context"
  "time"
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "gopkg.in/kataras/iris.v6/middleware/logger"
)

func hello(ctx *iris.Context) {
  ctx.HTML(iris.StatusOK, "Hello!! from " + ctx.Path())
}

func main() {
  app := iris.New()

  // output startup banner and error logs on os.Stdout
  app.Adapt(iris.DevLogger())

  // set the router, you can choose gorillamux too
  app.Adapt(httprouter.New())

  // request logger
  customLogger := logger.New(logger.Config{
    // Status displays status code
    Status: true,
    // IP displays request's remote address
    IP: true,
    // Method displays the http method
    Method: true,
    // Path displays the request path
    Path: true,
  })
  app.Use(customLogger)

  // 404
  errorLogger := logger.New()
  app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
    errorLogger.Serve(ctx)
    ctx.HTML(iris.StatusNotFound, "<h1>404</h1>")
  })

  // hello
  app.Get("/", hello)

  app.Adapt(iris.EventPolicy{
    // Interrupt Event means when control+C pressed on terminal.
    Interrupted: func(*iris.Framework) {
      // shut down gracefully, but wait 5 seconds the maximum before closed
      ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
      app.Shutdown(ctx)
    },
  })

  // execute the http server
  app.Listen(":8080")
}
