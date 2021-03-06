package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const Usage = "Usage: picam2s3 <bucket> <url>"

func fetchImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("saveImage failed. status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func uploadToS3(bucket, key string, body []byte) error {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
		ACL:    aws.String("private"),
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}

	bucket := os.Args[1]
	url := os.Args[2]

	body, err := fetchImage(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	key := fmt.Sprintf("%d.jpg", time.Now().Unix())

	if err := uploadToS3(bucket, key, body); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Snapshot was uploaded to s3://" + bucket + "/" + key)
}
