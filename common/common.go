package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"testcache/cacheengine"
	"time"
)

func GetLiveData(requestURL string, rqHash uint64) *cacheengine.CachedItem {
	var lCachedData *cacheengine.CachedItem = nil
	log.Printf("Request=%s\n", requestURL)
	var lCD cacheengine.CachedItem
	if res, errRs := http.Get(requestURL); errRs == nil {
		lCD.LastDate = time.Now()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			lCD.Error = err
		} else {
			lCD.Data = body
		}
		res.Close = true
	} else {
		log.Println("error making http Request")
	}
	lCachedData = &lCD
	return lCachedData
}
