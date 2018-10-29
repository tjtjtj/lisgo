package command

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CreateBucket
// DeleteBucket

// ListObject
// GetObject
// PutObject
// DeleteObject
// GetObjectAcl

const BucketsDir = "../testbuckets"
const Separator = "/"

func ListBuckets() ([]os.FileInfo, error) {
	var ret []os.FileInfo
	files, err := ioutil.ReadDir(BucketsDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			ret = append(ret, f)
			fmt.Println(f.Name())
		}
	}
	return ret, nil
}

func ListObjects(bucketname string, path string) ([]os.FileInfo, error) {
	var ret []os.FileInfo
	files, err := ioutil.ReadDir(BucketsDir + "/" + bucketname + "/" + path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			ret = append(ret, f)
			fmt.Println(f.Name())
		}
	}
	return ret, nil
}
