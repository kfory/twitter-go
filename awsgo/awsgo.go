package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Ctg aws.Config //aws = github.com/aws/aws-sdk-go-v2/aws
var err error

func InicializoAWS() {
	Ctx = context.TODO()                                                            //crea un context vacio
	Ctg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1")) //config = github.com/aws/aws-sdk-go-v2/config

	if err != nil {
		panic("Error al cargar la configuracion .aws/config" + err.Error())
	}

}
