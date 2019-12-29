package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	
	// Method: POST
	// Resource: http://localhost:8080/
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	
	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World"})
	})
	app.Run(iris.Addr(":8080"))
}
