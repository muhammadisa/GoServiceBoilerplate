package aliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/labstack/echo/v4"
)

// Repository interface
type Repository interface {
	GetObjects(bucketName string) (*oss.ListObjectsResult, error)
	GetBuckets() (*oss.ListBucketsResult, error)
	StoreObject(e echo.Context, bucketName string, tag string) ([]string, error)
	Delete(bucketName string, objectKey string) error
}
