package restaurantproducts

import (
	"context"
	"fmt"
	"mongodb/config"
	"mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

// func ProductContext() *mongo.Collection {
// 	client, _ := config.ConnectDataBase()
// 	return config.GetCollection(client, "inventory", "products")
// }

// func InsertProduct(product models.Product) (*mongo.InsertOneResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel() // Always remember to cancel the context to release resources
// 	result, err := ProductContext().InsertOne(ctx, product)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	fmt.Println(result)
// 	return result, nil
// }

// func InsertProductList(products []interface{}) (*mongo.InsertManyResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	result, err := ProductContext().InsertMany(ctx, products)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	fmt.Println(result)
// 	return result, nil
// }

func FindProducts() ([]*models.Restaurant, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.D{}

    client, err := config.ConnectDataBase()
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return nil, err
    }
    defer client.Disconnect(ctx)

    collection := config.GetCollection(client, "sample_restaurants", "restaurants")

    options := options.Find() // You can configure options here if needed
    result, err := collection.Find(ctx, filter, options)
    if err != nil {
        fmt.Println(err.Error())
        return nil, err
    }
    defer result.Close(ctx)
    count :=0
    var products []*models.Restaurant
    for result.Next(ctx) {
        product := &models.Restaurant{}
        err := result.Decode(product)
        if err != nil {
            fmt.Println(err.Error())
            return nil, err
        }
        fmt.Println(product)
        if count>1{break}else{
        products = append(products, product)
        }
        count++
    }

    if err := result.Err(); err != nil {
        fmt.Println(err.Error())
        return nil, err
    }

    return products, nil
}