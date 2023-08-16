// package services

// import (
// 	"context"
// 	"fmt"
// 	"mongodb/config"
// 	"mongodb/models"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"

// )

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

// func FindProducts() ([]*models.Product, error) {
//     ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//     defer cancel()

//     filter := bson.D{{"name", "Samsung"}}

//     client, err := config.ConnectDataBase()
//     if err != nil {
//         fmt.Println("Error connecting to the database:", err)
//         return nil, err
//     }
//     defer client.Disconnect(ctx)

//     collection := config.GetCollection(client, "inventory", "products")

//     options := options.Find() // You can configure options here if needed
//     result, err := collection.Find(ctx, filter, options)
//     if err != nil {
//         fmt.Println(err.Error())
//         return nil, err
//     }
//     defer result.Close(ctx)

//     var products []*models.Product
//     for result.Next(ctx) {
//         product := &models.Product{}
//         err := result.Decode(product)
//         if err != nil {
//             fmt.Println(err.Error())
//             return nil, err
//         }
//         products = append(products, product)
//     }

//     if err := result.Err(); err != nil {
//         fmt.Println(err.Error())
//         return nil, err
//     }

//     return products, nil
// }