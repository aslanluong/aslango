package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase() *mongo.Database  {
	ctx := context.Background()
	if client, err := mongo.NewClient(options.Client().ApplyURI("")); err != nil {
		fmt.Println(err)
	} else {
		if err := client.Connect(ctx); err != nil {
			fmt.Println(err)
		}
		return client.Database("Aslango")
	}
	return nil
}

func GetClient() *mongo.Client  {
	ctx := context.Background()
	if client, err := mongo.NewClient(options.Client().ApplyURI("")); err != nil {
		fmt.Println(err)
	} else {
		if err := client.Connect(ctx); err != nil {
			fmt.Println(err)
		}
		return client
	}
	return nil
}