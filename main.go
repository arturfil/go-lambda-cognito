package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/arturfil/go_lambdas/awsgo"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(
ctx context.Context,
event events.CognitoEventUserPoolsPostConfirmation,) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AwsInit()
	if !ValidParams() {
		fmt.Println("Error: should send 'SecretManager")
		err := errors.New("error: in params should send 'SecretManager'")
		return event, err
	}
	return events.CognitoEventUserPoolsPostConfirmation{}, nil
}

func ValidParams() bool {
	var hasParam bool
	_, hasParam = os.LookupEnv("SecretName")
	return hasParam
}
