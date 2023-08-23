package connections

import (
	"context"
	"fmt"
	"log"
	"mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "mongodb+srv://vasenth01:ved@college.yx7msef.mongodb.net/?retryWrites=true&w=majority"
const dbname = "college"
const collect = "profiles"

var collection *mongo.Collection

func Init() {

	clientOptions := options.Client().ApplyURI(connectionstring)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting")
		fmt.Println(err)
	}
	collection = client.Database(dbname).Collection(collect)

	fmt.Print("im done")
}
func Inserts(profile *models.Stuents) {

	inserted, err := collection.InsertOne(context.Background(), profile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted", inserted.InsertedID)
}

func Updates(ids string) {
	//provide the filter which one i need to filter
	//pass a flag called "set"
	//converting the ids into something that mongodb understands

	//primitive.ObjectIDFromHex(ids) //converts ids to mongodb uperstandable formate

	id, _ := primitive.ObjectIDFromHex(ids)

	// provid the filter
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"publishedYear": 1234, "author": "killers", "title": "The book"}}

	//bson.D and bosn.M difference is (M will give more clear result)
	//everying inside mongo is technically bson

	//now call the collections and update
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("modified successfully ", result.MatchedCount)
	//this will give the ouput of numbr of matches in results
}

func Deleteone(ids string) {

	id, _ := primitive.ObjectIDFromHex(ids)
	filter := bson.M{"_id": id}
	//in filter id is enough to delete

	result, _ := collection.DeleteOne(context.Background(), filter)
	fmt.Println("Delete ",result.DeletedCount)

}
