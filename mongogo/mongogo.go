package mongogo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	Date        time.Time `bson:"date"`
	Amount      float64   `bson:"amount"`
	Description string    `bson:"description"`
}

func connect() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("financeDB").Collection("transactions")
	return collection
}

func addTransaction(collection *mongo.Collection, transaction Transaction) {
	insertResult, err := collection.InsertOne(context.TODO(), transaction)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func retrieveTransactions(collection *mongo.Collection) {
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var elem Transaction
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Retrieved: %+v\n", elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	collection := connect()
	transaction := Transaction{
		Date:        time.Now(),
		Amount:      1234.56,
		Description: "Purchase at Store",
	}
	addTransaction(collection, transaction)
	retrieveTransactions(collection)
}
