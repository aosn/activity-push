#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o main .
../../../../../../bin/build-lambda-zip -o ../../main.zip main
rm main
