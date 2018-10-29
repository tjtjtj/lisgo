package server

import (
	"fmt"
	"io/ioutil"
)

// ListBuckets
func (s *Server) ListBuckets() ([]Bucket, error) {
	var ret []Bucket
	files, err := ioutil.ReadDir(s.BucketsDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name())
			b := Bucket{
				Name:         f.Name(),
				CreationDate: f.ModTime(),
			}
			ret = append(ret, b)
		}
	}
	return ret, nil
}

// ListObjects
// S3の挙動とはことなる
func (s *Server) ListObjects(bucketname string, path string) ([]Content, error) {
	var ret []Content
	files, err := ioutil.ReadDir(s.BucketsDir + "/" + bucketname + "/" + path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			c := Content{
				Key:          f.Name(),
				LastModified: f.ModTime(),
				Size:         f.Size(),
				StorageClass: "STANDARD",
				//TODO ETag string
				//TODO Owner Owner
			}
			ret = append(ret, c)
			//fmt.Println(f.Name())
		}
	}
	return ret, nil
}
