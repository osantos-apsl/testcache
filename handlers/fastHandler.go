package handlers

import (
	"log"
	"strings"
	"testcache/cacheengine"
	"testcache/common"
	"time"

	"github.com/valyala/fasthttp"
)

func FastHandler(ctx *fasthttp.RequestCtx) {
	pth := string(ctx.Path())
	if !strings.Contains(pth, "/q=") {
		ctx.Error("Not valid", 404)
		return
	}

	pCmd := strings.TrimPrefix(pth, "/")
	lCmdArr := strings.Split(pCmd, "=")
	if len(lCmdArr) != 2 {
		ctx.Error("Not valid", 404)
		return
	}

	var lCachedData *cacheengine.CachedItem
	switch lCmdArr[0] {
	case "q":
		start := time.Now()
		var lCached bool = true
		//test if is cached present.
		//requestURL := "http://localhost:8090/echo?q=" + pResult.Param
		requestURL := "https://postman-echo.com/get?q=" + lCmdArr[1]

		rqHash := common.FvnHash([]byte(requestURL))

		lCacheManager := cacheengine.GetCacheManager()
		lCachedData = lCacheManager.Get(&rqHash)
		if lCachedData == nil {
			lCachedData = common.GetLiveData(requestURL, rqHash)
			lCacheManager.Set(&rqHash, lCachedData)
			lCached = false
		}
		elapsed := time.Since(start)
		log.Printf("Itm Cached: %t Id: %d Addr: %p Time: %s", lCached, rqHash, lCachedData, elapsed)
	case "s":
		// instruccion for set data content.
		lCachedData = new(cacheengine.CachedItem)
	default:
		lCachedData = new(cacheengine.CachedItem)
	}

	if len(lCachedData.Data) <= 0 {
		ctx.Error("No data item", 404)
	}
	// set some headers and status code first
	ctx.SetContentType("text/html")
	ctx.SetStatusCode(fasthttp.StatusOK)

	ctx.SetBody(lCachedData.Data)

}
