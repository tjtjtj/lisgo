package handler

import (
	"net/http"

	"fmt"

	"github.com/labstack/gommon/log"
)

func Top(w http.ResponseWriter, r *http.Request) {
	log.Debugf("request url: %s", r.URL)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Top %v", r.Body)
}

func Articles(w http.ResponseWriter, r *http.Request) {
	log.Debugf("request url: %s", r.URL)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Articles %v", r.Body)
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Debugf("request url: %s", r.URL)
	fmt.Printf("Get %s %v", r.URL.Path, r.Body)

	// listbuckets
	//if r.URL.Path == "/" {
	//}

	// listobjects
	// getobjectacl
	// getobject

	w.WriteHeader(http.StatusOK)
	//fmt.Printf("Get %s %v", r.URL.Path, r.Body)
}
