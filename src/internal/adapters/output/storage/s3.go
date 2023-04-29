package adapter

import (
	"bytes"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	log        *log.Logger
	bucket     string
	client     *s3.S3
	downloader *s3manager.Downloader
}

func NewStorage(b string, l *log.Logger) *Storage {
	sess := session.Must(session.NewSession())
	client_s3 := s3.New(sess)
	return &Storage{
		log:        l,
		bucket:     b,
		client:     client_s3,
		downloader: s3manager.NewDownloader(sess),
	}
}

func (s Storage) GetListFiles(address string) ([]string, error) {
	result, err := s.client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.bucket),
	})
	if err != nil {
		s.log.Errorf("failed to list objects %s", err.Error())
		return nil, err
	}
	objectNames := make([]string, 0)
	for _, obj := range result.Contents {
		objectNames = append(objectNames, *obj.Key)
	}
	s.log.Infof("Documents list: %d", len(objectNames))
	return objectNames, nil
}

func (s Storage) GetFile(address string) (io.ReadCloser, error) {
	resp, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(address),
	})
	if err != nil {
		s.log.Errorf("failed to getting objects %s", err.Error())
		return nil, err
	}
	s.log.Infof("Document: %d", *resp.ContentLength)
	return resp.Body, nil
}

func (s Storage) PutFile(address string, doc *bytes.Buffer) (string, error) {
	fileContent := doc.Bytes()
	params := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(address),
		Body:   bytes.NewReader(fileContent), // Contenido del archivo
	}

	result, err := s.client.PutObject(params)
	if err != nil {
		return "", err
	}
	s.log.Infof("Doc saved %s", *result.ETag)
	url, err := s.getURL(address)
	if err != nil {
		return "", err
	}
	s.log.Infof("URL: %s", *result.ETag)
	return url, nil
}

func (s Storage) getURL(address string) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(address),
	})
	urlStr, err := req.Presign(5 * time.Hour)
	if err != nil {
		s.log.Errorf("Failed to sign request %s", err.Error())
		return "", err
	}

	s.log.Infof("The URL is %s", urlStr)
	return urlStr, nil
}
