package aliyun

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
)

// FileTagValidator query tag checker for file tag
func FileTagValidator(tag string) (string, error) {

	validTags := []string{
		"sinature",
		"profile",
		"report",
		"proof",
		"bug",
		"misc",
	}

	var usedTag string

	for _, value := range validTags {
		if tag == value {
			usedTag = value
			break
		}
	}

	if usedTag == "" {
		return usedTag, fmt.Errorf("%s tag not found", tag)
	}

	return usedTag, nil

}

// CreateAliyunOSSClient connect to aliyun from aliyun oss credential from env
func CreateAliyunOSSClient() (*oss.Client, string, error) {

	// Loading .env file
	var err error
	err = godotenv.Load()

	// Checking error for loading .env file
	if err != nil {
		return &oss.Client{}, "", fmt.Errorf("Error getting env, not coming through %v", err)
	}
	fmt.Println("We are getting the env values")

	endpoint := os.Getenv("PUBLIC_ENDPOINT_ALIYUNOSS")
	publicEndpoint := fmt.Sprintf("http://%s", endpoint)
	accessKeyID := os.Getenv("ACCESSKEYID_ALIYUNOSS")
	accessKey := os.Getenv("ACCESSKEY_ALIYUNOSS")

	client, err := oss.New(
		publicEndpoint,
		accessKeyID,
		accessKey,
	)

	if err != nil {
		return &oss.Client{}, "", err
	}

	return client, endpoint, nil

}
