package upload

import (
	"context"
	"fmt"
	"log"
	cfg "lokasani/app/drivers/config"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func ConfigCloud() *s3.Client {
	accountId := cfg.CloudAccount()
	accessKeyId := cfg.CloudKeyId()
	accessKeySecret := cfg.CloudKeySecret()

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com/lokasani", accountId),
		}, nil
	})

	fmt.Println(accountId)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return client
}

func UploadFile(file *multipart.FileHeader, client *s3.Client) string {

	bucketName := cfg.CloudBucket()
	fmt.Println(file)
	src, _ := file.Open()
	randomKey := uuid.New().String()
	input := &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &randomKey,
		Body:        src,
		ContentType: aws.String("image/png"),
	}

	_, err := client.PutObject(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	publiccostume := fmt.Sprintf("https://pub-0e2f2cf14fb64c84930c7dc185185ee2.r2.dev/%s", randomKey)
	return publiccostume

}
