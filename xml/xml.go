package xml

import (
	"encoding/xml"
	"time"
)

/*
<?xml version="1.0" encoding="UTF-8"?>
<ListAllMyBucketsResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Owner>
    <ID>8050</ID>
	<DisplayName>buckettest00002</DisplayName>
  </Owner>
  <Buckets>
    <Bucket><Name>readonly</Name><CreationDate>2016-04-27T23:41:33+0900</CreationDate></Bucket>
    <Bucket><Name>sandbox</Name><CreationDate>2016-04-27T23:44:52+0900</CreationDate></Bucket>
  </Buckets>
</ListAllMyBucketsResult>
*/
type ListAllMyBucketsResult struct {
	XMLName xml.Name `xml:"ListAllMyBucketsResult"`
	Xmlns   string   `xml:"xmlns,attr"`
	Owner   Owner
	Buckets struct {
		Bucket []Bucket
	}
	Bs []Bucket
}

type Owner struct {
	ID          string
	DisplayName string
}

type Bucket struct {
	Name         string
	CreationDate time.Time
}

/*
<?xml version="1.0" encoding="UTF-8"?>
<ListBucketResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Name>readonly</Name>
  <Prefix></Prefix>
  <Marker></Marker>
  <MaxKeys>1000</MaxKeys>
  <IsTruncated>false</IsTruncated>
  <Contents>
    <Key>ObjectStorage_API.pdf</Key>
    <LastModified>2016-04-27T23:41:27+0900</LastModified>
    <Size>412821</Size>
    <ETag>"3a7e05728edf101c044bceff755090e7"</ETag>
    <StorageClass>STANDARD</StorageClass>
    <Owner>
      <ID>8050</ID>
      <DisplayName>buckettest00002</DisplayName>
    </Owner>
  </Contents>
</ListBucketResult>
*/
type ListBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Xmlns       string   `xml:"xmlns,attr"`
	Name        string
	Prefix      string
	Marker      string
	MaxKeys     int
	IsTruncated bool
	Contents    Content
}

type Content struct {
	Key          string
	LastModified time.Time
	Size         int
	ETag         string
	StorageClass string
	Owner        Owner
}
