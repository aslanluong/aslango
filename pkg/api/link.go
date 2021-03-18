package api

import (
	"aslango/pkg/models"
	"aslango/pkg/platforms/mongodb"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
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

func GetShortLink(originalLink string) (*models.Link, error) {
	link := models.Link{}
	col := mongodb.GetDatabase().Collection("links")
	if err := col.FindOne(context.TODO(), bson.M{"original_link": originalLink}).Decode(&link); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &link, nil
}

func GenerateShortLink(originalLink string) (*models.Link, error) {
	if existingLink, _ := GetShortLink(originalLink); existingLink != nil {
		return existingLink, nil
	}

	linkId := primitive.NewObjectID()
	buf := make([]byte, 16)
	rand.Read(buf)
	shortLink := hex.EncodeToString(buf)[:9]

	for i := 1; i <= 5; i++ {
		if i == 5 {
			return nil, errors.New("get short link fail")
		}
		if _, err := GetOriginalLink(shortLink); err != nil {
			break
		}
		rand.Read(buf)
		shortLink = hex.EncodeToString(buf)[:9]
	}

	link := models.Link{
		LinkId:       linkId,
		OriginalLink: originalLink,
		ShortLink:    shortLink,
		Active:       0,
		CreateAt:     linkId.Timestamp(),
		UpdateAt:     linkId.Timestamp(),
	}

	col := mongodb.GetDatabase().Collection("links")
	if _, err := col.InsertOne(context.TODO(), link); err != nil {
		fmt.Println("Insert fail")
		return nil, err
	}
	return &link, nil
}
