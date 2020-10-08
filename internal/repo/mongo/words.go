package mongo

import (
	"fmt"
	"context"
	"time"
	d "github.com/logologics/kunren-be/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

// StoreWord bla
func (mongo *Mongo) StoreWord(w d.Word) (mp.ObjectID, error) {
	words := mongo.kunrenDB.Collection("words")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// check exists
	opts := mopt.Find()
	opts.SetProjection(bson.M{"_id": 0})
	fRes := words.FindOne(ctx, bson.D{mp.E{Key: "key", Value: w.Key}})
	if fRes.Err() == nil {
		var word d.Word
		fRes.Decode(&word)
		fmt.Printf("New id %v", word.ID)
		return word.ID, nil
	}

	fmt.Printf("Find err %v", fRes.Err())

	w.ID = mp.NewObjectID()
	iRes, err := words.InsertOne(ctx, w)
	if err != nil {
		fmt.Printf("Insert err %v", err)

		return mp.ObjectID{}, err
	}

	fmt.Printf("New id %v", iRes.InsertedID.(mp.ObjectID))
	return iRes.InsertedID.(mp.ObjectID), nil
}

// LoadWord bla
func (mongo *Mongo) LoadWord(id mp.ObjectID) (d.Word, error) {
	return d.Word{}, nil
}

// UpdateWord bla
func (mongo *Mongo) UpdateWord(w d.Word) error {
	return nil
}

// DeleteWord bla
func (mongo *Mongo) DeleteWord(id mp.ObjectID) error {
	return nil
}

// ListWords lists all words in the dict
func (mongo *Mongo) ListWords() ([]d.Word, error) {
	return nil, nil
}
