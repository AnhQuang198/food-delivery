package configs

import (
	"food-delivery/components/uploadprovider"
	"os"
)

func GetS3Connection() uploadprovider.UploadProvider {
	s3BucketName := os.Getenv("s3BucketName")
	s3Region := os.Getenv("s3Region")
	s3APIKey := os.Getenv("s3APIKey")
	s3SecretKey := os.Getenv("s3SecretKey")
	s3Domain := os.Getenv("s3Domain")
	return uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
}
