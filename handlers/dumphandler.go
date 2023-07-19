package handlers

import (
	"fmt"
	"net/http"
	"testcache/cacheengine"
	"unsafe"
)

func DumpHandler(pResponse http.ResponseWriter, pRequest *http.Request) {
	cManager := cacheengine.GetCacheManager()
	lAllData := cManager.Dump()
	var lTotalSize int = 0
	var lItms int = 0
	for i := range lAllData {
		lTam := len((*lAllData[i]).Data) + int(unsafe.Sizeof(*lAllData[i]))
		fmt.Fprintf(pResponse, "Itm Id: %d Addr: %p Time: %s Size: %d\n", (*lAllData[i]).HashID, lAllData[i], (*lAllData[i]).LastDate, lTam)
		lTotalSize += lTam
		lItms += 1
	}
	fmt.Fprintf(pResponse, "Total Resume / Items: %d Cache Size: %d Kb\n", lItms, lTotalSize/1024)
}
