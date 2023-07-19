package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testcache/cacheengine"
	"testcache/common"
	"time"
)

type CmdResult struct {
	Cmd   string
	Param string
	Error error
}

// Process data, from cached or live.
func cmdProcess(pResult *CmdResult, pCmdProcess chan cacheengine.CachedItem) {
	var lCachedData *cacheengine.CachedItem
	switch pResult.Cmd {
	case "q":
		start := time.Now()
		var lCached bool = true
		//test if is cached present.
		//requestURL := "http://localhost:8090/echo?q=" + pResult.Param
		requestURL := "https://postman-echo.com/get?q=" + pResult.Param

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
	pCmdProcess <- *lCachedData
}

// Extract commnad to send (goRutine)
func cmdParse(pCmd string, pChCmmd chan CmdResult) {
	pCmd = strings.TrimPrefix(pCmd, "/")
	lCmdArr := strings.Split(pCmd, "=")
	lRes := new(CmdResult)
	if len(lCmdArr) != 2 {
		lRes.Error = errors.New("Command error")
	} else {
		lRes.Cmd = lCmdArr[0]
		lRes.Param = lCmdArr[1]
	}
	pChCmmd <- *lRes
}

func ChannelHandler(pResponse http.ResponseWriter, pRequest *http.Request) {
	if !strings.Contains(pRequest.URL.Path, "/q=") {
		http.NotFound(pResponse, pRequest)
		return
	}

	ch_Cmd := make(chan CmdResult, 1)
	ch_Proc := make(chan cacheengine.CachedItem, 1)
	ctx := pRequest.Context()
	go cmdParse(pRequest.URL.Path, ch_Cmd)

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(pResponse, err.Error(), internalError)
	case lCmdRes := <-ch_Cmd:
		if lCmdRes.Error != nil {
			internalError := http.StatusInternalServerError
			http.Error(pResponse, lCmdRes.Error.Error(), internalError)
			return
		}
		go cmdProcess(&lCmdRes, ch_Proc)
	}

	select {
	case lCachedData := <-ch_Proc:
		{
			if len(lCachedData.Data) > 0 {
				pResponse.Write(lCachedData.Data)
			} else {
				fmt.Fprintln(pResponse, lCachedData.Error)
			}

		}
	}
}
