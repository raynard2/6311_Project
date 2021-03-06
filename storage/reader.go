package storage

import (
	"6311_Project/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * Save
 * Save is used to save a record in the mongoStore
 */

func (d *mongoconn) ReaderSave(user interface{}, Collection string) (models.Reader, error) {
	var output models.Reader
	_result, err := d.mongoDb.Collection(Collection).InsertOne(context.Background(), user)

	if err != nil {
		return output, err
	}

	err = d.mongoDb.Collection(Collection).FindOne(nil, bson.M{"_id": _result.InsertedID}).Decode(&output)

	if err != nil {
		return output, err
	}
	return output, nil
}



func (d *mongoconn) FindManyHistory(fields, projection, sort map[string]interface{}, limit, skip int64, results interface{}) error {
	ops := options.Find()
	if limit > 0 {
		ops.Limit = &limit
	}
	if skip > 0 {
		ops.Skip = &skip
	}
	if projection != nil {
		ops.Projection = projection
	}
	if sort != nil {
		ops.Sort = sort
	}


	cursor, err :=  d.mongoDb.Collection("readerEventType").Find(nil, fields, ops)

	//cursor, err := d.mongoDb.Collection("books").Find(nil, fields, ops)
	if err != nil {
		return err
	}

	var output []map[string]interface{}
	for cursor.Next(nil) {
		var item map[string]interface{}
		_ = cursor.Decode(&item)
		output = append(output, item)
		//fmt.Println(output)
	}

	if b, e := json.Marshal(output); e == nil {
		_ = json.Unmarshal(b, &results)

		//fmt.Println(results)
	} else {
		return e
	}
	return nil
}
