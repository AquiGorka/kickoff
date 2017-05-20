package main 

import (
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func hello(ctx *iris.Context) {
    ctx.Writef("Hello from %s", ctx.Path())
}

func notFound(ctx *iris.Context) {
  ctx.HTML(iris.StatusNotFound, "<h1>404</h1>")
}

func main() {
  app := iris.New()
  // output startup banner and error logs on os.Stdout
  //app.Adapt(iris.DevLogger())
  // set the router, you can choose gorillamux too
  app.Adapt(httprouter.New())

  app.Get("/", hello)
/*
  app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
    ctx.HTML(iris.StatusNotFound, "<h1>Custom not found handler </h1>")
  })
*/
/*
  app.HandleFunc("GET", "/", func(ctx *iris.Context) {
    ctx.Writef("hello world\n")
  })
*/
/*
  app.Get("/hi", func(ctx *iris.Context) {
    ctx.HTML(iris.StatusOK, " <h1>hi, I just exist in order to see if the server is closed</h1>")
  })
*/
/*
  app.Adapt(iris.EventPolicy{
    // Interrupt Event means when control+C pressed on terminal.
    Interrupted: func(*iris.Framework) {
      // shut down gracefully, but wait 5 seconds the maximum before closed
      ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
      app.Shutdown(ctx)
    },
  })
*/
  app.Listen(":8080")
}
