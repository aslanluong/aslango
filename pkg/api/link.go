package api

import (
	"aslango/pkg/models"
	"aslango/pkg/platforms/mongodb"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetOriginalLink(shortLink string) (*string, error) {
	link := models.Link{}
	col := mongodb.GetDatabase().Collection("links")
	fmt.Println(shortLink)
	if err := col.FindOne(context.TODO(), bson.M{"short_link": shortLink}).Decode(&link); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &link.OriginalLink, nil
}