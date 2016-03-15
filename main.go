package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadFile(filename string) {
	s3Filename := filename + ".gz"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	// gzip file contents
	reader, writer := io.Pipe()
	go func() {
		gw := gzip.NewWriter(writer)
		io.Copy(gw, file)

		file.Close()
		gw.Close()
		writer.Close()
	}()

	// Upload gzipped contents
	uploader := s3manager.NewUploader(session.New(&aws.Config{
		Credentials: getCredentials(),
		Region:      aws.String(awsRegion)}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   reader,
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Filename),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	log.Println("Successfully uploaded to", result.Location)
}

func main() {
	uploadFile("test.txt")
}
