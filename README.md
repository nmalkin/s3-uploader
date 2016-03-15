S3 Uploader
===========

This utility, when executed, uploads the specified files to an S3 bucket. For
ease of use, all of the configurations are hard-coded in the program itself.
The files will be gzipped as part of the upload process.

Configuration
-------------

For this program to work, you'll need to specify the following settings in `config.go`:

- the filenames of the files to be uploaded
- the name of the S3 bucket to upload the files too. (This should be created ahead of time.)
- the AWS region where the bucket is located
- your AWS user's key ID and secret access key

Building
--------

To build the program, run:

    go build -o upload *.go

To build for another platform:

    GOOS=windows GOARCH=386 go build -o upload.exe *.go

