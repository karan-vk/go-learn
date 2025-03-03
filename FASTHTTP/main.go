package main

import (
	"github.com/valyala/fasthttp"
	"fmt"
)

type MyHandler struct {
	foobar string
	l int
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "hello %+v",h)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

// pass bound struct method to fasthttp
func main() {
	myHandler := &MyHandler{
		foobar: "foobar",
		l:5,
	}
	fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)

	// pass plain function to fasthttp
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}
