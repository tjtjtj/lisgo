package server

import (
	"io/ioutil"
	"os"
)

// ListBuckets
func (s *Server) ListBuckets() ([]os.FileInfo, error) {
	var ret []os.FileInfo
	files, err := ioutil.ReadDir(s.BucketsDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			ret = append(ret, f)
		}
	}
	return ret, nil
}

// ListObjects
// S3の挙動とはことなる
func (s *Server) ListObjects(bucketname string, path string) ([]os.FileInfo, error) {
	var ret []os.FileInfo
	files, err := ioutil.ReadDir(s.BucketsDir + "/" + bucketname + "/" + path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			ret = append(ret, f)
			//fmt.Println(f.Name())
		}
	}
	return ret, nil
}
