package utils

import (
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func JsonResponse(ctx *fasthttp.RequestCtx, status int, data interface{}) {
	ctx.Response.Header.Set("Content-type", "application/json")
	ctx.Response.SetStatusCode(status)
	response, err := json.Marshal(data)
	if err != nil {
		ctx.Error("Internal server error: ", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Response.SetBody(response)
}

func JsonError(ctx *fasthttp.RequestCtx, status int, message string) {
	JsonResponse(ctx, status, map[string]string{
		"error": message,
	})
}
