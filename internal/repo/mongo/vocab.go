package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"time"
)

// StoreVocab bla
func (mongo *Mongo) StoreVocab(v d.Vocab) (mp.ObjectID, error) {
	words := mongo.kunrenDB.Collection("vocab")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// check exists
	opts := mopt.Find()
	opts.SetProjection(bson.M{"_id": 0})

	fRes := words.FindOne(ctx, bson.D{mp.E{Key: "_id", Value: v.ID}})
	if fRes != nil {
		return v.ID, nil
	}

	iRes, err := words.InsertOne(ctx, v)
	if err != nil {
		return mp.ObjectID{}, err
	}

	return iRes.InsertedID.(mp.ObjectID), nil

}

// LoadVocab blax
func (mongo *Mongo) LoadVocab(id mp.ObjectID) (d.Vocab, error) {
	return d.Vocab{}, nil
}

// UpdateVocab bla
func (mongo *Mongo) UpdateVocab(v d.Vocab) error {
	return nil
}

// UpsertVocab bla
func (mongo *Mongo) UpsertVocab(v d.Vocab) (mp.ObjectID, error) {
	return mp.ObjectID{}, nil
}

// DeleteVocab bla
func (mongo *Mongo) DeleteVocab(id mp.ObjectID) error {
	return nil
}

// ListVocab lists all Vocab ofr the user
func (mongo *Mongo) ListVocab(u d.User, st r.SortType) ([]d.Vocab, error) {
	return nil, nil
}
