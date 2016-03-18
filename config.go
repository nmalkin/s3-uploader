package main

import (
	"os"
	"path/filepath"
)

var config = struct {
	AWSRegion          string
	Bucket             string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	DirectoryToUpload  string
}{
	"us-west-2",
	"MY-BUCKET",
	"",
	"",
	filepath.Join(os.Getenv("TEMP"), "upload"),
}
