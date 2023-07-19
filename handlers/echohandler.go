package handlers

import (
	"fmt"
	"net/http"
)

func EchoHandler(pResponse http.ResponseWriter, pRequest *http.Request) {
	fmt.Fprintf(pResponse, "%s", pRequest.URL.RawQuery)
}
