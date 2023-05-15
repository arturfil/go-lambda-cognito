package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/arturfil/go_lambdas/awsgo"
	"github.com/arturfil/go_lambdas/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
    var secretData models.SecretRDSJson
    fmt.Println("> Asking for secret", secretName)

    svc := secretsmanager.NewFromConfig(awsgo.Cfg)
    key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
        SecretId: aws.String(secretName),
    })
    if err != nil {
        fmt.Println(err.Error())
        return secretData, err 
    }
    
    json.Unmarshal([]byte(*key.SecretString), &secretData)
    fmt.Println("Reading Secret Ok")

    return secretData, nil
}
