package server

import "time"

type Server struct {
	BucketsDir string
}

type Bucket struct {
	Name         string
	CreationDate time.Time
}

type Content struct {
	Key          string
	LastModified time.Time
	Size         int64
	ETag         string
	StorageClass string
	Owner        Owner
}

type Owner struct {
	ID          string
	DisplayName string
}
