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

func main() {
	bucket := "tor-ux"
	filename := "test.txt"
	s3Key := "test.txt.gz"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	// Not required, but you could zip the file before uploading it
	// using io.Pipe read/writer to stream gzip'd file contents.
	reader, writer := io.Pipe()
	go func() {
		gw := gzip.NewWriter(writer)
		io.Copy(gw, file)

		file.Close()
		gw.Close()
		writer.Close()
	}()
	uploader := s3manager.NewUploader(session.New(&aws.Config{
		Credentials: getCredentials(),
		Region:      aws.String("us-west-2")}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   reader,
		Bucket: aws.String(bucket),
		Key:    aws.String(s3Key),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	log.Println("Successfully uploaded to", result.Location)
}
