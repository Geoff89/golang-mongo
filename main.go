package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:pass@localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// create a goroutone to defer the closure
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	usersCollection := client.Database("testing").Collection("users")
	/*
		// insert a single document in a collection ussing bson.D object
		user := bson.D{{"fullname", "User 1"}, {"age", 30}}
		// insert bson object using insertOne()
		result, err := usersCollection.InsertOne(context.TODO(), user)
		if err != nil {
			panic(err)
		}
		// display the id of the newly created objet
		fmt.Println(result.InsertedID)

		users := []interface{}{
			bson.D{{"fullName", "User 2"}, {"age", 25}},
			bson.D{{"fullName", "User 3"}, {"age", 20}},
			bson.D{{"fullName", "User 4"}, {"age", 28}},
		}
		// insert the bson object slice using InsertMany()
		results, err := usersCollection.InsertMany(context.TODO(), users)
		// check for errors in the insertion
		if err != nil {
			panic(err)
		}
		// display the ids of the newly inserted objects
		fmt.Println(results.InsertedIDs)
	*/
	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// retrieve all the documents that match the filter
	cursor, err := usersCollection.Find(context.TODO(), filter)
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results1 []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results1); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results1 {
		fmt.Println(result)
	}

	// retrieving the first document that matches the filter
	var result2 bson.M
	// check for errors in the finding
	if err = usersCollection.FindOne(context.TODO(), filter).Decode(&result2); err != nil {
		panic(err)
	}

	// display the document retrieved
	fmt.Println("displaying the first result from the search filter")
	fmt.Println(result2)

	// update a single document with a specified ObjectID using UpdateByID()
	// insert a new document to the collection
	user := bson.D{{"fullName", "User 5"}, {"age", 22}}
	insertResult, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	// create the update query for the client
	update := bson.D{
		{"$set",
			bson.D{
				{"fullName", "User V"},
			},
		},
		{"$inc",
			bson.D{
				{"age", 1},
			},
		},
	}

	// execute the UpdateByID() function with the filter and update query
	result, err := usersCollection.UpdateByID(context.TODO(), insertResult.InsertedID, update)
	// check for errors in the updating
	if err != nil {
		panic(err)
	}
	// display the number of documents updated
	fmt.Println("Number of documents updated:", result.ModifiedCount)

	// update single and multiple documents with a specified filter using UpdateOne() and UpdateMany()
	// create a search filer
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// create the update query
	update := bson.D{
		{"$set",
			bson.D{
				{"age", 40},
			},
		},
	}

	// execute the UpdateOne() function to update the first matching document
	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	// check for errors in the updating
	if err != nil {
		panic(err)
	}
	// display the number of documents updated
	fmt.Println("Number of documents updated:", result.ModifiedCount)

	// execute the UpdateMany() function to update all matching first document
	results, err := usersCollection.UpdateMany(context.TODO(), filter, update)
	// check for errors in the updating
	if err != nil {
		panic(err)
	}
	// display the number of documents updated
	fmt.Println("Number of documents updated:", results.ModifiedCount)

	// replace the fields of a single document with ReplaceOne()
	// create a search filer
	filter := bson.D{{"fullName", "User 1"}}

	// create the replacement data
	replacement := bson.D{
		{"firstName", "John"},
		{"lastName", "Doe"},
		{"age", 30},
		{"emailAddress", "johndoe@email.com"},
	}

	// execute the ReplaceOne() function to replace the fields
	result, err := usersCollection.ReplaceOne(context.TODO(), filter, replacement)
	// check for errors in the replacing
	if err != nil {
		panic(err)
	}
	// display the number of documents updated
	fmt.Println("Number of documents updated:", result.ModifiedCount)

	// delete single and multiple documents with a specified filter using DeleteOne() and DeleteMany()
	// create a search filter
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// delete the first document that match the filter
	result, err := usersCollection.DeleteOne(context.TODO(), filter)
	// check for errors in the deleting
	if err != nil {
		panic(err)
	}
	// display the number of documents deleted
	fmt.Println("deleting the first result from the search filter")
	fmt.Println("Number of documents deleted:", result.DeletedCount)

	// delete every document that match the filter
	results, err := usersCollection.DeleteMany(context.TODO(), filter)
	// check for errors in the deleting
	if err != nil {
		panic(err)
	}
	// display the number of documents deleted
	fmt.Println("deleting every result from the search filter")
	fmt.Println("Number of documents deleted:", results.DeletedCount)

}
