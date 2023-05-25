package db

import (
	"context"
	"fmt"
	"intent-service/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IntentDb struct {
	db         *mongo.Client
	collection string
	database   string
}

func NewIntentDb(
	db *mongo.Client,
	collection string,
	database string,
) (*IntentDb, error) {
	fdb := db.Database(database)

	err := fdb.CreateCollection(context.TODO(), collection, nil)

	if err != nil && err.Error() != fmt.Sprintf("(NamespaceExists) Collection %s.%s already exists.", database, collection) {
		return nil, err
	}

	return &IntentDb{db: db, collection: collection, database: database}, nil
}

func (itp *IntentDb) Save(intent entities.IntentInterface) (entities.IntentInterface, error) {
	coll := itp.db.Database(itp.database).Collection(itp.collection)

	_, err := coll.InsertOne(context.Background(), intent)

	if err != nil {
		return nil, err
	}

	return intent, nil
}

func (itp *IntentDb) Get(id string) (entities.IntentInterface, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	coll := itp.db.Database(itp.database).Collection(itp.collection)
	filter := bson.D{{Key: "_id", Value: objectID}}
	var result entities.Intent

	if err := coll.FindOne(context.Background(), filter).Decode(&result); err != nil {
		return nil, err
	}

	resultParsedDto := entities.IntentDto{
		Itens:   result.Itens,
		User_id: result.UserID,
	}

	resultParsed, err := entities.NewIntent(resultParsedDto)

	if err != nil {
		return nil, err
	}

	return resultParsed, nil
}

func (itp *IntentDb) GetAll() ([]entities.IntentInterface, error) {
	coll := itp.db.Database(itp.database).Collection(itp.collection)
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []entities.Intent

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	var resultsParsed []entities.IntentInterface
	for _, result := range results {
		resultParsedDto := entities.IntentDto{
			Itens:   result.Itens,
			User_id: result.UserID,
		}

		resultParsed, err := entities.NewIntent(resultParsedDto)

		if err != nil {
			return nil, err
		}

		resultsParsed = append(resultsParsed, resultParsed)
	}

	return resultsParsed, nil
}
