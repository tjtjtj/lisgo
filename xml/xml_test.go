package xml

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjtjtj/lisgo/server"
)

func Test_BucketsResult(t *testing.T) {
	owner := server.Owner{
		ID:          "testid",
		DisplayName: "testname",
	}
	b1 := server.Bucket{
		Name: "b1name",
	}
	b2 := server.Bucket{
		Name: "b2name",
	}
	var list []server.Bucket
	list = append(list, b1, b2)
	result := BucketsResult(owner, list)

	byteXML, err := xml.MarshalIndent(result, "", `   `)
	assert.NoError(t, err)
	s := xml.Header + string(byteXML)
	fmt.Printf("xml:%s\n", s)
	assert.Equal(t, "", s)
}

func Test_ObjectsResult(t *testing.T) {
	c := server.Content{
		Key: "readme.txt",
		//LastModified:
		Size:         12345,
		ETag:         "\"aaaaaaaaaa\"",
		StorageClass: "STANDARD",
	}
	c.Owner.ID = "testid"
	c.Owner.DisplayName = "testname"
	result := ObjectsResult(c)

	byteXML, err := xml.MarshalIndent(result, "", `   `)
	assert.NoError(t, err)
	s := xml.Header + string(byteXML)
	fmt.Printf("xml:%s\n", s)
	assert.Equal(t, "", s)
}

func TestXML_ListAllMyBucketsResult(t *testing.T) {
	var result ListAllMyBucketsResult
	result.Xmlns = "http://s3.amazonaws.com/doc/2006-03-01/"
	result.Owner.ID = "asdf"
	result.Owner.DisplayName = "asdfasdf"
	b1 := server.Bucket{
		Name: "b1name",
	}
	b2 := server.Bucket{
		Name: "b2name",
	}
	result.Buckets.Bucket = append(result.Buckets.Bucket, b1, b2)

	byteXML, err := xml.MarshalIndent(result, "", `   `)
	assert.NoError(t, err)
	s := xml.Header + string(byteXML)
	fmt.Printf("xml:%s\n", s)
	assert.Equal(t, "", s)
}

func TestXML_ListBucketResult(t *testing.T) {
	var result ListBucketResult
	result.Xmlns = "http://s3.amazonaws.com/doc/2006-03-01/"
	result.Name = "example"
	result.Prefix = "pre"
	result.Marker = "mark"
	result.MaxKeys = 1000
	result.IsTruncated = false
	result.Contents.Key = "readme.txt"
	//result.Contents.LastModified =
	result.Contents.Size = 12345
	result.Contents.ETag = "\"aaaaaaaaaa\""
	result.Contents.StorageClass = "STANDARD"
	result.Contents.Owner.ID = "asdf"
	result.Contents.Owner.DisplayName = "asdfasdf"

	byteXML, err := xml.MarshalIndent(result, "", `   `)
	assert.NoError(t, err)
	s := xml.Header + string(byteXML)
	fmt.Printf("xml:%s\n", s)
	assert.Equal(t, "", s)
}
