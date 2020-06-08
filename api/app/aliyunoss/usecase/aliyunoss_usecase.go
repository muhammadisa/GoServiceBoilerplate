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
func NewAliyunOSSUsecase(aoUsecase aliyunoss.Repository) aliyunoss.Usecase {
	return &aliyunOSSUsecase{
		aliyunOSSRepository: aoUsecase,
	}
}

func (aoUsecase aliyunOSSUsecase) GetBuckets() (*oss.ListBucketsResult, error) {
	buckets, err := aoUsecase.aliyunOSSRepository.GetBuckets()
	if err != nil {
		return nil, err
	}
	return buckets, nil
}

func (aoUsecase aliyunOSSUsecase) GetObjects(bucketName string) (*oss.ListObjectsResult, error) {
	aliyunObjects, err := aoUsecase.aliyunOSSRepository.GetObjects(bucketName)
	if err != nil {
		return nil, err
	}
	return aliyunObjects, nil
}

func (aoUsecase aliyunOSSUsecase) StoreObject(e echo.Context, bucketName string, tag string) (string, error) {
	publicEndpoint, err := aoUsecase.aliyunOSSRepository.StoreObject(e, bucketName, tag)
	if err != nil {
		return "", err
	}
	return publicEndpoint, nil
}

func (aoUsecase aliyunOSSUsecase) Delete(bucketName string, objectKey string) error {
	return aoUsecase.aliyunOSSRepository.Delete(bucketName, objectKey)
}
