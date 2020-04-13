package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-boilerplate/api/app/aliyunoss"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/aliyun"
)

type aliyunOSSInteractorRepo struct{}

// NewAliyunOSSInteractorRepo function
func NewAliyunOSSInteractorRepo() aliyunoss.Repository {
	return &aliyunOSSInteractorRepo{}
}

func (i *aliyunOSSInteractorRepo) GetObjects(bucketName string) (*oss.ListObjectsResult, error) {
	client, _, err := aliyun.CreateAliyunOSSClient()
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	lsRes, err := bucket.ListObjects()
	if err != nil {
		return nil, err
	}
	return &lsRes, nil
}

func (i *aliyunOSSInteractorRepo) GetBuckets() (*oss.ListBucketsResult, error) {
	client, _, err := aliyun.CreateAliyunOSSClient()
	if err != nil {
		return nil, err
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		return nil, err
	}
	return &lsRes, nil
}

func (i *aliyunOSSInteractorRepo) StoreObject(e echo.Context, bucketName string, tag string) (string, error) {
	client, publicEndpoint, err := aliyun.CreateAliyunOSSClient()
	if err != nil {
		return "", err
	}
	fileHeader, err := e.FormFile("file")
	if err != nil {
		return "", errors.New("Problem with file which you choose")
	}
	f, err := fileHeader.Open()
	if err != nil {
		return "", errors.New("Problem with file which you choose")
	}
	fileTimestamp := time.Now().Format("20060102-150405")
	fileType := strings.Split(fileHeader.Filename, ".")
	if len(fileType) != 2 {
		return "", errors.New("Unformatted file doesn't allowed to upload")
	}
	objectKey := fmt.Sprintf("heliosstorage-tag-%s-%s.%s", tag, fileTimestamp, fileType[1])

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(objectKey, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.%s/%s", bucketName, publicEndpoint, objectKey), nil
}

func (i *aliyunOSSInteractorRepo) Delete(bucketName string, objectKey string) error {
	client, _, err := aliyun.CreateAliyunOSSClient()
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(objectKey)
	if err != nil {
		return err
	}
	return nil
}
