package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct{
	LinkId primitive.ObjectID `bson:"_id" json:"id"`
	OriginalLink string `bson:"original_link" json:"original_link"`
	ShortLink string `bson:"short_link" json:"short_link"`
	Active int `bson:"active" json:"active"`
	CreateAt time.Time `bson:"create_at" json:"create_at"`
	UpdateAt time.Time `bson:"update_at" json:"update_at"`
}

