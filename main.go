package main

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// uploadFile uploads the file named "filename" from the directory "directory"
// to S3. It compresses the file using gzip, adds the extension ".gz", and puts
// it at the top level of the bucket.
//
// If a file with this filename already exists, it will be overwritten.
func uploadFile(directory string, filename string) {
	path := filepath.Join(directory, filename)
	s3Filename := filename + ".gz"

	file, err := os.Open(path)
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
		Credentials: credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, ""),
		Region:      aws.String(config.AWSRegion)}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   reader,
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(s3Filename),
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	log.Println("Successfully uploaded", filename, "to", result.Location)
}

func main() {
	files, _ := ioutil.ReadDir(config.DirectoryToUpload)
	for _, file := range files {
		uploadFile(config.DirectoryToUpload, file.Name())
	}
}
