package middlewares

import (
	"strings"

	"github.com/hassannmoussaa/bookery/pkg/appCtx"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/valyala/fasthttp"
)

func ParseQueryParams(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	fn := func(requestCtx *fasthttp.RequestCtx) {
		ctx := appCtx.Get(requestCtx)
		ctx.QueryParams.Search = getSearchFromRequest(requestCtx)
		ctx.QueryParams.Set("search", helpers.BytesToString(requestCtx.QueryArgs().Peek("search")))
		ctx.QueryParams.Page = getPageNumberFromRequest(requestCtx)
		ctx.QueryParams.Set("page", helpers.BytesToString(requestCtx.QueryArgs().Peek("page")))
		ctx.QueryParams.SinceID = getSinceIDFromRequest(requestCtx)
		ctx.QueryParams.Set("since_id", helpers.BytesToString(requestCtx.QueryArgs().Peek("since_id")))
		ctx.QueryParams.Count = getCountFromRequest(requestCtx)
		ctx.QueryParams.Set("count", helpers.BytesToString(requestCtx.QueryArgs().Peek("count")))
		ctx.QueryParams.From = getFromFromRequest(requestCtx)
		ctx.QueryParams.Set("from", helpers.BytesToString(requestCtx.QueryArgs().Peek("from")))
		if next != nil {
			next(requestCtx)
		}
	}
	return fn
}

func getSearchFromRequest(requestCtx *fasthttp.RequestCtx) string {
	if search := helpers.BytesToString(requestCtx.QueryArgs().Peek("search")); search != "" {
		return strings.TrimSpace(search)
	}
	return ""
}
func getPageNumberFromRequest(requestCtx *fasthttp.RequestCtx) int32 {
	pageNumber := requestCtx.QueryArgs().GetUintOrZero("page")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	return int32(pageNumber)
}
func getSinceIDFromRequest(requestCtx *fasthttp.RequestCtx) int32 {
	sinceID := requestCtx.QueryArgs().GetUintOrZero("since_id")
	if sinceID > 0 {
		return int32(sinceID)
	}
	return 0
}
func getCountFromRequest(requestCtx *fasthttp.RequestCtx) int32 {
	count := requestCtx.QueryArgs().GetUintOrZero("count")
	if count > 0 {
		return int32(count)
	}
	return 0
}
func getFromFromRequest(requestCtx *fasthttp.RequestCtx) int32 {
	from := requestCtx.QueryArgs().GetUintOrZero("from")
	if from > 0 {
		return int32(from)
	}
	return 0
}
