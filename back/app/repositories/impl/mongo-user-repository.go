package impl

import (
	"context"

	"github.com/iryoda/price-guru/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) MongoUserRepository {
	return MongoUserRepository{Db: db, Collection: db.Collection("users")}
}

func (ur MongoUserRepository) Create(user *entities.User) (*entities.User, error) {
	result, err := ur.Collection.InsertOne(context.TODO(), user)

	if err != nil {
		return user, err
	}

	user.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return user, nil
}

func (ur MongoUserRepository) FindById(id string) (*entities.User, error) {
	var user entities.User

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = ur.Collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur MongoUserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := ur.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
