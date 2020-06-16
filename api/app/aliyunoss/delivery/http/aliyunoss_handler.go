package http

import (
	"net/http"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/labstack/echo/v4"
	"github.com/muhammadisa/go-service-boilerplate/api/app/aliyunoss"
	"github.com/muhammadisa/go-service-boilerplate/api/response"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/message"
	uuid "github.com/satori/go.uuid"
)

// AliyunOSSHandler struct
type AliyunOSSHandler struct {
	aOUsecase aliyunoss.Usecase
}

// NewAliyunOSSHandler initialize endpoints
func NewAliyunOSSHandler(e *echo.Group, aOu aliyunoss.Usecase) {
	handler := &AliyunOSSHandler{
		aOUsecase: aOu,
	}
	e.GET("/aliyunoss/buckets/", handler.GetBuckets)
	e.GET("/aliyunoss/objects/", handler.GetObjects)
	e.POST("/aliyunoss/objects/", handler.Store)
	e.DELETE("/aliyunoss/rollback/objects/:bucket/:key/", handler.Delete)
}

// GetBuckets Get aliyun oss buckets data
func (aoHandler *AliyunOSSHandler) GetBuckets(c echo.Context) error {
	buckets, err := aoHandler.aOUsecase.GetBuckets()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message: message.GenerateMessage(uuid.Nil, "GET", oss.ListBucketsResult{}, false) +
				" Cannot retrieve buckets",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uuid.Nil, "GET", oss.ListBucketsResult{}, true),
		Data:       buckets,
	})
}

// GetObjects Get aliyun oss objects from bucket data
func (aoHandler *AliyunOSSHandler) GetObjects(c echo.Context) error {
	bucket := c.QueryParam("bucket")

	aliyunObjects, err := aoHandler.aOUsecase.GetObjects(bucket)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message: message.GenerateMessage(uuid.Nil, "GET", oss.ListObjectsResult{}, false) +
				" Cannot retrieve objects",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uuid.Nil, "GET", oss.ListObjectsResult{}, true),
		Data:       aliyunObjects,
	})
}

// Store file aliyun oss
func (aoHandler *AliyunOSSHandler) Store(c echo.Context) error {
	publicEndpoint, err := aoHandler.aOUsecase.StoreObject(
		c,
		c.FormValue("bucket"),
		c.FormValue("tag"),
	)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message: message.GenerateMessage(uuid.Nil, "POST", c.File, false) +
				" Cannot put objects",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uuid.Nil, "POST", c.File, true),
		Data: map[string]string{
			"file_url": publicEndpoint[0],
			"object":   publicEndpoint[1],
		},
	})
}

// Delete rolling back file from aliyun oss
func (aoHandler *AliyunOSSHandler) Delete(c echo.Context) error {
	bucket, key := c.Param("bucket"), c.Param("key")
	err := aoHandler.aOUsecase.Delete(bucket, key)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message: message.GenerateMessage(uuid.Nil, "DELETE", c.File, false) +
				" Cannot delete objects",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    message.GenerateMessage(uuid.Nil, "DELETE", c.File, true),
		Data: map[string]string{
			"source":         bucket,
			"deleted_object": key,
		},
	})
}
