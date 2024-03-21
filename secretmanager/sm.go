package secretmanager

import (
	"encoding/json" //permite vtrabajar con toda la codificacion de json
	"fmt"
	"twitter-go/awsgo"
	"twitter-go/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("> Pido secreto " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Ctg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> lectura de secret ok " + secretName)
	return datosSecret, nil
}
