package api

import (
	"aslango/pkg/models"
	"aslango/pkg/platforms/mongodb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOriginalLink(shortLink string) (*models.Link, error) {
	link := models.Link{}
	col := mongodb.GetDatabase().Collection("links")
	fmt.Println(shortLink)
	if err := col.FindOne(context.TODO(), bson.M{"short_link": shortLink}).Decode(&link); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &link, nil
}

func UpdateLinkActive(linkId primitive.ObjectID) {
	col := mongodb.GetDatabase().Collection("links")
	update := bson.M{"$currentDate": bson.M{"update_at": true}, "$inc": bson.M{"active": 1}}
	if _, err := col.UpdateByID(context.TODO(), linkId, update); err != nil {
		fmt.Println("Update fail!")
	}
}
