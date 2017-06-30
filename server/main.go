package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func httpHandler(w http.ResponseWriter, req *http.Request) {
	r, err := http.Get("http://api:8781/")
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			uTime, err := strconv.ParseInt(string(body), 10, 64)
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				gTime := time.Unix(uTime, 0)
				sTime := gTime.String()
				w.Write([]byte(sTime))
			}
		}
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(httpHandler))

	log.Fatal(http.ListenAndServe(":8780", nil))
}
