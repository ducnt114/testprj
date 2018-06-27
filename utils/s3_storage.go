package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"path/filepath"
)

type s3_storage struct {
	Session *session.Session
}

// Create new AWS S3 client --
// token is optional if id and secret is provide both
func NewS3Client(region, id, secret, token string) (*s3_storage, error) {
	conf := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(id, secret, token),
	}
	sess, err := session.NewSession(conf)
	if err != nil {
		log.Println("Error when create aws session, detail: ", err)
		return nil, err
	}
	return &s3_storage{Session: sess}, nil
}

func (s *s3_storage) UploadFile(fileName, bucket string) (string, error) {
	svc := s3manager.NewUploader(s.Session)

	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Failed to open file", fileName, err)
		return "", err
	}
	defer file.Close()

	log.Println("Uploading file to S3...")
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filepath.Base(fileName)),
		Body:   file,
	})
	if err != nil {
		log.Println("Error when upload file to s3, detail: ", err)
		return "", err
	}

	log.Printf("Successfully uploaded %s to %s\n", fileName, result.Location)
	return result.Location, nil
}
