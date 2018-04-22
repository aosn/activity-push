// Copyright 2018 mikan.

package main

import (
	"bytes"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Upload(name string, data []byte) {
	bucket := "chart.aosn.ws"
	service := s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApNortheast1RegionID),
	})))
	var expires = time.Now().Add(time.Hour * 24 * 6) // 6 days later
	_, err := service.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(name),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
		ACL:         aws.String(s3.BucketCannedACLPublicRead),
		Expires:     &expires,
	})
	if err != nil {
		panic(err)
	}
}
