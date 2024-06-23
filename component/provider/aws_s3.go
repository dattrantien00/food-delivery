package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"food-delivery/common"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName, region, apiKey, secret, domain string) *s3Provider {
	prod := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
		// session: session,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			apiKey, secret, "",
		),
	})
	if err != nil {
		log.Fatalln(err)
	}
	prod.session = s3Session

	return prod
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(provider.bucketName),
		Key : aws.String(dst),
		ACL: aws.String("private"),
		ContentType: aws.String(fileType),
		Body: fileBytes,

	})
	if err != nil{
		return nil,err
	}
	return &common.Image{
		Url: fmt.Sprintf("%s/%s", provider.domain, dst),
	},nil
}
