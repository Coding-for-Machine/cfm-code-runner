package handlers

import (
	"runner/internal/utils"

	"github.com/valyala/fasthttp"
)

func HomeAPI(ctx *fasthttp.RequestCtx) {
	data := map[string]interface{}{
		"page":    "Home page",
		"version": "v1",
	}
	utils.JsonResponse(
		ctx,
		fasthttp.StatusOK,
		data,
	)
}
