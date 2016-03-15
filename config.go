package main

var config = struct {
	AWSRegion          string
	Bucket             string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	FilesToUpload      []string
}{
	"us-west-2",
	"MY-BUCKET",
	"",
	"",
	[]string{"README.md", "test.txt"},
}
