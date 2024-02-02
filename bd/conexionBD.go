package bd

import (
	"context"
	"fmt"
	"twitter-go/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConectarBD(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string) // para convertir un context en string
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, passwd, host)

	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion Exitosa con la BD...")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil

}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil //condicion //puede ser true o false
}
