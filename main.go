package main

import (
	"context"
	"os"
	"strings"
	"twitter-go/awsgo"
	"twitter-go/bd"
	"twitter-go/handlers"
	"twitter-go/models"
	"twitter-go/secretmanager"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse //variable de tipo (*events.APIGatewayProxyResponse)

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. deben incluir 'SecretName','BucketName','UrlPrefix'",
			Headers: map[string]string{
				"content-type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("secretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de secret" + err.Error(),
			Headers: map[string]string{
				"content-type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitter"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)                          //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)          //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)          //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)      //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)              //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)      //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTsign)        //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)                  //establecer variables
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName")) //establecer variables

	//chuequeo conecion ala BD o conecto la BD
	err = bd.ConectarBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando la BD" + err.Error(),
			Headers: map[string]string{
				"content-type": "application/json",
			},
		}
		return res, nil
	}

	respAPI := handlers.Manejadores(awsgo.Ctx, request)

	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"content-type": "application/json",
			},
		}
		return res, nil
	} else {
		return respAPI.CustomResp, nil
	}
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro { //si no trae parametro (devuelve un false)
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro { //si no trae parametro (devuelve un false)
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro { //si no trae parametro (devuelve un false)
		return traeParametro
	}

	return traeParametro
}
