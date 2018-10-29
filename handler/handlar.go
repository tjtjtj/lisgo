package handler

import (
	"net/http"

	"github.com/tjtjtj/lisgo/server"

	"fmt"

	"github.com/labstack/gommon/log"
)

func Get(w http.ResponseWriter, r *http.Request) {
	log.Debugf("request url: %s", r.URL)
	fmt.Printf("Get %s %v", r.URL.Path, r.Body)

	srv := server.Server{BucketsDir: "./buckets"}

	// listbuckets
	if r.URL.Path == "/" {
		respbody, err := listobjects(&srv)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respbody))
		return
	}

	// listobjects

	// getobjectacl
	// getobject

	w.WriteHeader(http.StatusOK)
	//fmt.Printf("Get %s %v", r.URL.Path, r.Body)
}
