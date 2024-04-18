package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDotEnv() {
	godotenv.Load()
}

func InitMongo() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return mongoClient
}

func InitAmqp() *amqp.Queue {
	uri := os.Getenv("AMQP_URI")
	queue := os.Getenv("AMQP_QUEUE")

	conn, err := amqp.Dial(uri)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // argumentse

	)

	if err != nil {
		panic(err)
	}

	return &q
}

func main() {
	InitDotEnv()

	c := InitMongo()
	q := InitAmqp()

	d := c.Database("price-guru").Collection("watchers")

	d.Find(context.Background(), bson.M{})
}
