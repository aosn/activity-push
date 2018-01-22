// Copyright 2018 mikan.

package main

import "github.com/aws/aws-lambda-go/lambda"

// main executes handler.
func main() {
	lambda.Start(CollectAndPush)
}
