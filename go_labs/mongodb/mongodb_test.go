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
	ID int64  `json:"id" bson:"_id"`
	A  string `json:"a"`
	B  string `json:"b"`
}

var (
	testDataArr = []testData{
		{
			ID: 0,
			A:  "a_0",
			B:  "b_0",
		},
		{
			ID: 1,
			A:  "a_1",
			B:  "b_1",
		},
		{
			ID: 2,
			A:  "a_2",
			B:  "b_2",
		},
		{
			ID: 3,
			A:  "a_3",
			B:  "b_3",
		},
		{
			ID: 4,
			A:  "a_4",
			B:  "b_4",
		},
		{
			ID: 5,
			A:  "a_5",
			B:  "b_5",
		},
		{
			ID: 6,
			A:  "a_6",
			B:  "b_6",
		},
		{
			ID: 7,
			A:  "a_7",
			B:  "b_7",
		},
		{
			ID: 8,
			A:  "a_8",
			B:  "b_8",
		},
	}
)

func Test_CRUD(t *testing.T) {
	Test_InsertOne(t)

	Test_InsertMany(t)

	Test_FindOne(t)

	Test_Find(t)

	Test_ReplaceOne(t)

	Test_UpdateOne(t)

	Test_UpdateByID(t)

	Test_UpdateMany(t)

	Test_DeleteOne(t)

	Test_DeleteMany(t)
}

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

func Test_FindOne(t *testing.T) {
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

func Test_Find(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	filter := bson.M{}

	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		t.Error(err)
	}

	var resultArr []testData

	err = cur.All(context.TODO(), &resultArr)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(resultArr); i++ {
		if resultArr[i].ID != testDataArr[i].ID || resultArr[i].A != testDataArr[i].A || resultArr[i].B != testDataArr[i].B {
			t.Error("Invalid result")
		}
	}
}

func Test_ReplaceOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	type R_testData struct {
		ID int64  `json:"id" bson:"_id"`
		C  string `json:"c"`
	}

	targetDataID := testDataArr[0].ID

	filter := bson.M{"_id": targetDataID}
	replacement := R_testData{
		ID: 0,
		C:  "c_0",
	}

	res, err := coll.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		t.Error(err)
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 || res.UpsertedCount != 0 || res.UpsertedID != nil {
		t.Error("Invalid result")
	}
}

func Test_UpdateOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	targetData := testDataArr[1]
	targetData.A += "_update"
	targetData.B += "_update"

	filter := bson.M{"a": testDataArr[1].A}
	update := bson.M{"$set": targetData}

	res, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 || res.UpsertedCount != 0 || res.UpsertedID != nil {
		t.Error("Invalid result")
	}
}

func Test_UpdateByID(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	targetData := testDataArr[2]
	targetData.A += "updatebyid"
	targetData.B += "updatebyid"

	filter := targetData.ID
	update := bson.M{"$set": targetData}

	res, err := coll.UpdateByID(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 || res.UpsertedCount != 0 || res.UpsertedID != nil {
		t.Error("Invalid result")
	}
}

func Test_UpdateMany(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	filter := bson.M{"_id": bson.M{"$gte": testDataArr[3].ID}}
	update := bson.M{"$set": bson.M{"a": "a", "b": "b"}}

	res, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		t.Error(err)
	}

	updatedCount := len(testDataArr[3:])

	if res.MatchedCount != int64(updatedCount) || res.ModifiedCount != int64(updatedCount) || res.UpsertedCount != 0 || res.UpsertedID != nil {
		t.Error("Invalid result")
	}
}

func Test_DeleteOne(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	targetData := testDataArr[0]

	filter := bson.M{"_id": targetData.ID}

	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Error(err)
	}

	if res.DeletedCount != 1 {
		t.Error("Invalid result")
	}
}

func Test_DeleteMany(t *testing.T) {
	coll, close := GetCollection()
	defer close()

	filter := bson.M{"_id": bson.M{"$gte": testDataArr[1].ID}}

	res, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		t.Error(err)
	}

	if res.DeletedCount != int64(len(testDataArr[1:])) {
		t.Error("Invalid result")
	}
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

/// Comparison Query Operators
// $eq	: Matches values that are equal to a specified value.
// $gt	: Matches values that are greater than a specified value.
// $gte	: Matches values that are greater than or equal to a specified value.
// $in	: Matches any of the values specified in an array.
// $lt	: Matches values that are less than a specified value.
// $lte	: Matches values that are less than or equal to a specified value.
// $ne	: Matches all values that are not equal to a specified value.
// $nin	: Matches none of the values specified in an array.
// ex) filter := bson.D{{"rating", bson.D{{"$lt", 7}}}}

/// Logical Query Operators
// $and	: Joins query clauses with a logical AND returns all documents that match the conditions of both clauses.
// $not	: Inverts the effect of a query expression and returns documents that do not match the query expression.
// $nor	: Joins query clauses with a logical NOR returns all documents that fail to match both clauses.
// $or	: Joins query clauses with a logical OR returns all documents that match the conditions of either clause.
// ex)	filter := bson.D{
//			{"$and",
//			   bson.A{
//				  bson.D{{"rating", bson.D{{"$gt", 7}}}},
//				  bson.D{{"rating", bson.D{{"$lte", 10}}}},
//			   },
//			},
//		}

/// Element Query Operators
// $exists	: Matches documents that have the specified field.
// $type	: Selects documents if a field is of the specified type.
// ex) filter := bson.D{{"vendor", bson.D{{"$exists", false}}}}

/// Evaluation Query Operators
// $expr		 : Allows use of aggregation expressions within the query language.
// $jsonSchema	 : Validate documents against the given JSON Schema.
// $mod			 : Performs a modulo operation on the value of a field and selects documents with a specified result.
// $regex		 : Selects documents where values match a specified regular expression.
// $text		 : Performs text search.
// $where		 : Matches documents that satisfy a JavaScript expression.
// ex) filter := bson.D{{"type", bson.D{{"$regex", "^E"}}}}
