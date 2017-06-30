package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func httpHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
}

func main() {
	http.Handle("/", http.HandlerFunc(httpHandler))

	log.Fatal(http.ListenAndServe(":8781", nil))
}
