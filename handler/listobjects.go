package handler

import (
	"encoding/xml"

	"github.com/tjtjtj/lisgo/server"
	"github.com/tjtjtj/lisgo/xmlconverter"
)

func listobjects(srv *server.Server) (string, error) {
	owner := server.Owner{
		ID:          "testid",
		DisplayName: "testname",
	}
	buckets, err := srv.ListBuckets()
	if err != nil {
		return "", err
	}
	result := xmlconverter.BucketsResult(owner, buckets)
	byteXML, _ := result.MarshalXml()
	if err != nil {
		return "", err
	}
	return xml.Header + string(byteXML), nil
}
