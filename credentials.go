package main

import "github.com/aws/aws-sdk-go/aws/credentials"

func getCredentials() *credentials.Credentials {
	aws_access_key_id := ""
	aws_secret_access_key := ""
	token := ""

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	return creds
}