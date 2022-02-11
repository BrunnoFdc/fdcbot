package database

import (
	"context"
	"fdcteam-bot/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var currentClient *mongo.Client

func DatabaseClient() *mongo.Client {
	return currentClient
}

func Connect() (err error) {
	log.Info("Iniciando conexão com o banco de dados")

	connectionOptions := options.Client().ApplyURI(config.MongodbConnectionUrl)

	var client *mongo.Client

	client, err = mongo.NewClient(connectionOptions)

	if err != nil {
		log.Error("Erro ao criar cliente do MongoDB", err)
		return
	}

	err = client.Connect(context.TODO())

	if err != nil {
		log.Error("Erro ao conectar com o banco de dados", err)
		return
	}

	currentClient = client

	return
}
