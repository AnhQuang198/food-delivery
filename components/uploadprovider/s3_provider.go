package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"food-delivery/commons"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	//create session and verify credential on aws
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             // Token can be ignore
	})

	if err != nil {
		log.Fatalln(err)
	}

	provider.session = s3Session

	return provider
}

//push object to aws-s3
func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*commons.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	img := &commons.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
