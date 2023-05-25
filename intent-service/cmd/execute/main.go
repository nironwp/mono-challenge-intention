package main

import (
	"context"
	"fmt"
	db "intent-service/adapters/db/mongo"
	"intent-service/domain/entities"
	"log"
	"os"

	graphserver "intent-service/cmd/graph-server"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	database_name := os.Getenv("MONGODB_DATABASE")

	if database_name == "" {
		log.Fatal("You must set your 'MONGODB_DATABASE' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("You must set your 'PORT' environmental variable.")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	dbmongo, err := db.NewIntentDb(client, "intents", database_name)

	if err != nil {
		log.Fatal(err)
	}

	service := entities.IntentService{
		Persistence: dbmongo,
	}

	graphServer := graphserver.NewGraphServer(port, service)

	graphServer.Start()
}
