package handler

import (
	"encoding/xml"
	"net/http"

	"github.com/tjtjtj/lisgo/server"
	"github.com/tjtjtj/lisgo/xmlconverter"

	"fmt"

	"github.com/labstack/gommon/log"
)

func Get(w http.ResponseWriter, r *http.Request) {
	log.Debugf("request url: %s", r.URL)
	fmt.Printf("Get %s %v", r.URL.Path, r.Body)

	srv := server.Server{BucketsDir: "./buckets"}

	// listbuckets
	if r.URL.Path == "/" {
		owner := server.Owner{
			ID:          "testid",
			DisplayName: "testname",
		}
		buckets, _ := srv.ListBuckets()
		result := xmlconverter.BucketsResult(owner, buckets)
		byteXML, _ := result.MarshalXml()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(xml.Header))
		w.Write(byteXML)
		return
	}

	// listobjects
	// getobjectacl
	// getobject

	w.WriteHeader(http.StatusOK)
	//fmt.Printf("Get %s %v", r.URL.Path, r.Body)
}
