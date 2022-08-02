package mongodb

import (
	"context"
	"fmt"
	"log"
	"testing"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type info struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func GetCollection() (*mongo.Collection, func()) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", "root", "example", "127.0.0.1", "27017")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	disconnectDB := func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}

	return client.Database("insertDB").Collection("info"), disconnectDB
}

func Test_InsertOne(t *testing.T) {
	coll, closeFunc := GetCollection()
	defer closeFunc()

	data := info{
		ID:    "0",
		Title: "title",
		Text:  "text",
	}

	res, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.InsertedID)
}

func Test_InsertMany(t *testing.T) {
	coll, closeFunc := GetCollection()
	defer closeFunc()

	infoArr := []info{
		{
			ID:    "1",
			Title: "title1",
			Text:  "text1",
		},
		{
			ID:    "2",
			Title: "title2",
			Text:  "text2",
		},
		{
			ID:    "3",
			Title: "title3",
			Text:  "text3",
		},
	}

	doc := make([]interface{}, 0)
	for _, info := range infoArr {
		doc = append(doc, info)
	}

	result, err := coll.InsertMany(context.TODO(), doc)
	if err != nil {
		t.Error(err)
	}

	t.Log(result.InsertedIDs...)
}


func Test_DeleteOne(t *testing.T) {
	coll, closeFunc := GetCollection()
	defer closeFunc()

	filter := info{
		Title: "title",
	}

	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.DeletedCount)
}