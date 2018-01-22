// Copyright 2018 mikan.

package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// handler handles lambda context
func handler(ctx context.Context, config events.ConfigEvent) (string, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return "", errors.New("cannot find LambdaContext")
	}
	log.Printf("InvokedFunctionArn = %s\n", lc.InvokedFunctionArn)
	log.Printf("config version is %v\n", config.Version)
	log.Printf("\"%s\" executes on \"%s\".\n", os.Getenv("LAMBDA_TASK_ROOT"), os.Getenv("AWS_EXECUTION_ENV"))
	CollectAndPush()
	return "function finished", nil
}

// main executes handler.
func main() {
	lambda.Start(handler)
}
