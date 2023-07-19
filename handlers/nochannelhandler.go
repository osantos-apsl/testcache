package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testcache/cacheengine"
	"testcache/common"
	"time"
)

func NoChannelHandler(pResponse http.ResponseWriter, pRequest *http.Request) {

	if !strings.Contains(pRequest.URL.Path, "/q=") {
		http.NotFound(pResponse, pRequest)
		return
	}

	pCmd := strings.TrimPrefix(pRequest.URL.Path, "/")
	lCmdArr := strings.Split(pCmd, "=")
	if len(lCmdArr) != 2 {
		internalError := http.StatusInternalServerError
		http.Error(pResponse, "Command error", internalError)
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

	if len(lCachedData.Data) > 0 {
		pResponse.WriteHeader(200)
		pResponse.Write(lCachedData.Data)
	} else {
		fmt.Fprintln(pResponse, lCachedData.Error)
	}

}
