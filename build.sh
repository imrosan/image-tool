#! /bin/bash

export GO111MODULE=on
go mod init github.com/imrosan/image-tool
go build -o main ./
