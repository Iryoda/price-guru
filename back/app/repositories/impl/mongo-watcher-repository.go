package impl

import (
	"context"

	"github.com/iryoda/price-guru/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoWatcherRepository struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func NewMongoWatcherRepository(db *mongo.Database) MongoWatcherRepository {
	return MongoWatcherRepository{Db: db, Collection: db.Collection("watchers")}
}

func (wr MongoWatcherRepository) Create(watcher *entities.Watcher) (*entities.Watcher, error) {
	result, err := wr.Collection.InsertOne(context.TODO(), watcher)

	if err != nil {
		return watcher, err
	}

	watcher.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return watcher, nil
}

func (wr MongoWatcherRepository) FindAllByUserId(id string) (*[]entities.Watcher, error) {
	var watchers []entities.Watcher

	cursor, err := wr.Collection.Find(context.TODO(), bson.M{"userId": id})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var watcher entities.Watcher
		cursor.Decode(&watcher)
		watchers = append(watchers, watcher)
	}

	return &watchers, nil
}

func (wr MongoWatcherRepository) FindById(id string) (entities.Watcher, error) {
	var watcher entities.Watcher

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return watcher, err
	}

	err = wr.Collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&watcher)

	if err != nil {
		return watcher, err
	}

	return watcher, nil
}

func (wr MongoWatcherRepository) DeleteById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = wr.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		return err
	}

	return nil
}
