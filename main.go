package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/arturfil/go_lambdas/awsgo"
	"github.com/arturfil/go_lambdas/db"
	"github.com/arturfil/go_lambdas/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.AwsInit()
	if !ValidParams() {
		fmt.Println("Error: should send 'SecretManager")
		err := errors.New("error: in params should send 'SecretManager'")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = ", data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = ", data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret", err.Error())
		return event, err
	}

	err = db.SignUp(data)
    if err != nil {
        fmt.Println("ERROR -> ", err.Error())
    }
	return event, err
}

func ValidParams() bool {
	var hasParam bool
	_, hasParam = os.LookupEnv("SecretName")
	return hasParam
}
