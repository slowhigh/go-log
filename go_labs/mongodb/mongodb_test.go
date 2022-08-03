package mongodb

import (
	"context"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testData struct {
	ID string `json:"id" bson:"_id"`
	A  string `json:"a"`
	B  string `json:"b"`
}

func Test_CRUD(t *testing.T) {
	Test_InsertOne(t)

	Test_InsertMany(t)

	Test_Find(t)
}

var (
	testDataArr = []testData{
		{
			ID: "0",
			A:  "a_0",
			B:  "b_0",
		},
		{
			ID: "1",
			A:  "a_1",
			B:  "b_1",
		},
		{
			ID: "2",
			A:  "a_2",
			B:  "b_2",
		},
		{
			ID: "3",
			A:  "a_3",
			B:  "b_3",
		},
		{
			ID: "4",
			A:  "a_4",
			B:  "b_4",
		},
		{
			ID: "5",
			A:  "a_5",
			B:  "b_5",
		},
		{
			ID: "6",
			A:  "a_6",
			B:  "b_6",
		},
		{
			ID: "7",
			A:  "a_7",
			B:  "b_7",
		},
		{
			ID: "8",
			A:  "a_8",
			B:  "b_8",
		},
	}
)

func Test_InsertOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	data := testDataArr[0]

	res, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		t.Error(err)
	}

	if res.InsertedID != data.ID {
		t.Error("Invalid result")
	}
}

func Test_InsertMany(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	dataArr := testDataArr[1:]

	documents := make([]interface{}, 0)
	for _, data := range dataArr {
		documents = append(documents, data)
	}

	res, err := coll.InsertMany(context.TODO(), documents)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(res.InsertedIDs); i++ {
		if res.InsertedIDs[i] != dataArr[i].ID {
			t.Error("Invalid result")
		}
	}
}

func Test_Find(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	data := testDataArr[0]

	filter := bson.M{"_id": data.ID}

	res := coll.FindOne(context.TODO(), filter)

	if err := res.Err(); err != nil {
		t.Error(err)
	}

	var foundData testData

	if err := res.Decode(&foundData); err != nil {
		t.Error(err)
	}

	if foundData.ID != data.ID || foundData.A != data.A || foundData.B != data.B {
		t.Error("Invalid result")
	}
}

// => 여기까지 수정함

func Test_DeleteOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	filter := bson.M{"title": "title"} // filter := bson.D{{"title","title"}}

	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.DeletedCount)
}

func Test_DeleteMany(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	filter := bson.M{"_id": "1"} // filter := bson.D{{"title","title"}}

	res, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.DeletedCount)
}

func Test_UpdateOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	data := testData{
		ID:   "2",
		A:    "title_updateone",
		B: "text_updateone",
	}

	filter := bson.M{"_id": data.ID}
	update := bson.M{"$set": bson.M{"title": data.A, "text": data.B}}

	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.MatchedCount)
}

func Test_UpdateByID(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	data := testData{
		ID:   "3",
		A:    "title_updateone",
		B: "text_updateone",
	}

	filter := data.ID
	update := bson.M{"$set": bson.M{"title": data.A, "text": data.B}}

	res, err := coll.UpdateByID(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.MatchedCount)
}

func Test_UpdateMany(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	filter := bson.M{"text": "text"}
	update := bson.M{"$set": bson.M{"title": "title"}}

	res, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.MatchedCount)
}

func Test_ReplaceOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	data := testData{
		ID:   "2",
		A:    "title_updateone",
		B: "text_updateone",
	}

	filter := bson.M{"_id": data.ID}
	update := bson.M{"title": data.A}

	res, err := coll.ReplaceOne(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.MatchedCount)
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

	return client.Database("testDB").Collection("data"), disconnectDB
}
