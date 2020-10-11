package mongo

import (
	"context"
	"fmt"

	d "github.com/logologics/kunren-be/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

// StoreWord stores Work if it does not exist or if lemma fingerprint has changed
func (mongo *Mongo) StoreWord(word d.Word) (mp.ObjectID, error) {
	words := mongo.kunrenDB.Collection("words")
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	// # check exists and hash matches
	opts := mopt.FindOne()

	// ## find
	opts.SetProjection(bson.M{"lhash": 1})
	fRes := words.FindOne(ctx, bson.D{mp.E{Key: "key", Value: word.Key}}, opts)
	fResErr := fRes.Err()

	if fResErr != nil && fResErr != mlib.ErrNoDocuments {
		return mp.ObjectID{}, fRes.Err()
	}

	// ## calc hash
	h, hashErr := word.Lemma.Hash()
	if hashErr != nil {
		return mp.ObjectID{}, hashErr
	}

	// # if found 1
	if fResErr == nil {
		// ## decode
		var loadedWord d.Word
		err := fRes.Decode(&loadedWord)
		if err != nil {
			return mp.ObjectID{}, err
		}

		// ## return ID if hashes match
		if h == loadedWord.LemmaHash {
			return loadedWord.ID, nil
		}

		// ## update
		// update hash
		word.LemmaHash = h
		word.ID = loadedWord.ID
		uRes, err := words.ReplaceOne(ctx, bson.M{"key": word.Key}, word)
		if err != nil {
			return mp.ObjectID{}, err
		}
		if uRes.MatchedCount != 1 {
			return mp.ObjectID{}, fmt.Errorf("Nothing was updated for key %v", word.Key)
		}


		fmt.Printf("Updated id %v, old hash %v != new hash %v", word.ID, loadedWord.LemmaHash, h)
		return loadedWord.ID, nil
	}

	// ## otherwise insert
	word.ID = mp.NewObjectID()
	word.LemmaHash = h
	iRes, err := words.InsertOne(ctx, word)
	if err != nil {
		return mp.ObjectID{}, err
	}

	fmt.Printf("New word with id %v created", iRes.InsertedID.(mp.ObjectID))
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
