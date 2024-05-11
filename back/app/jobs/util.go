package jobs

import (
	"context"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
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

func InitAmqp() *amqp.Channel {
	uri := os.Getenv("RABBITMQ_URI")
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

	err = ch.ExchangeDeclare(
		os.Getenv("WATCHER_EXCHANGE"),
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	_, err = ch.QueueDeclare(
		os.Getenv("WATCHER_QUEUE"), // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // argumentse
	)

	if err != nil {
		panic(err)
	}

	return ch
}
