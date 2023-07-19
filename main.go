package main

import (
	"log"
	"net/http"
	"testcache/cacheengine"
	"testcache/handlers"
)

var cManager *cacheengine.CacheManager

func main() {
	
	// create logs
	log.SetPrefix("L: ")
	log.SetFlags(0)
	log.Println("init started")
	//  create cache layers
	cacheengine.GetCacheManager().RegisterLayer(cacheengine.CreateInMemory)
	//global listeners
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/dump", handlers.DumpHandler)
	http.HandleFunc("/", handlers.NoChannelHandler)
	http.ListenAndServe(":8090", nil)
}
