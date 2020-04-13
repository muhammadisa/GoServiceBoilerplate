package usecase

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-boilerplate/api/app/aliyunoss"
)

// aliyunOSSUsecase
type aliyunOSSUsecase struct {
	aliyunOSSRepository aliyunoss.Repository
}

// NewAliyunOSSUsecase function
func NewAliyunOSSUsecase(aO aliyunoss.Repository) aliyunoss.Usecase {
	return &aliyunOSSUsecase{
		aliyunOSSRepository: aO,
	}
}

func (aO aliyunOSSUsecase) GetBuckets() (*oss.ListBucketsResult, error) {
	buckets, err := aO.aliyunOSSRepository.GetBuckets()
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func (aO aliyunOSSUsecase) GetObjects(bucketName string) (*oss.ListObjectsResult, error) {
	aliyunObjects, err := aO.aliyunOSSRepository.GetObjects(bucketName)
	if err != nil {
		return nil, err
	}
	return aliyunObjects, nil
}

func (aO aliyunOSSUsecase) StoreObject(e echo.Context, bucketName string, tag string) (string, error) {
	publicEndpoint, err := aO.aliyunOSSRepository.StoreObject(e, bucketName, tag)
	if err != nil {
		return "", err
	}
	return publicEndpoint, nil
}

func (aO aliyunOSSUsecase) Delete(bucketName string, objectKey string) error {
	return aO.aliyunOSSRepository.Delete(bucketName, objectKey)
}
