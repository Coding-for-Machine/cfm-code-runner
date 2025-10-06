package routes

import (
	"runner/internal/handlers"

	"github.com/valyala/fasthttp"
)

func Router(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		handlers.HomeAPI(ctx)
	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBody([]byte(`{"error": "Not found"}`))
	}
}
