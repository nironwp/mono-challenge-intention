package db

import (
	"context"
	"intent-service/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IntentProductDb struct {
	db         *mongo.Client
	collection string
	database   string
}

func NewIntentProductDb(db *mongo.Client, collection string, database string) *IntentProductDb {
	return &IntentProductDb{db: db, collection: collection, database: database}
}

func (itp *IntentProductDb) GetAll() ([]entities.IntentProductInterface, error) {
	coll := itp.db.Database(itp.database).Collection(itp.collection)
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []entities.IntentProduct

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	var resultsParsed []entities.IntentProductInterface
	for _, result := range results {
		resultParsedDto := entities.NewIntentProductDto{
			Title:       result.Title,
			Price:       result.Price,
			Category:    result.Category,
			Image:       result.Image,
			ID:          result.ID,
			Description: result.Description,
		}

		resultParsedProduct, err := entities.NewIntentProduct(resultParsedDto)

		if err != nil {
			return nil, err
		}

		resultsParsed = append(resultsParsed, resultParsedProduct)
	}

	return resultsParsed, nil
}

func (itp *IntentProductDb) Get(id string) (entities.IntentProductInterface, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	coll := itp.db.Database(itp.database).Collection(itp.collection)
	filter := bson.D{{Key: "_id", Value: objectID}}
	var result entities.IntentProduct

	err = coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	resultParsedDto := entities.NewIntentProductDto{
		Title:       result.Title,
		Price:       result.Price,
		Category:    result.Category,
		Image:       result.Image,
		ID:          result.ID,
		Description: result.Description,
	}

	resultParsedProduct, err := entities.NewIntentProduct(resultParsedDto)

	if err != nil {
		return nil, err
	}

	return resultParsedProduct, nil
}

func (itp *IntentProductDb) Save(product entities.IntentProductInterface) (entities.IntentProductInterface, error) {
	coll := itp.db.Database(itp.database).Collection(itp.collection)

	_, err := coll.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
