package mongo

import (
	"context"
	"fmt"

	d "github.com/logologics/kunren-be/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
)

func (mongo *Mongo) wordsCollection() *mlib.Collection {
	return mongo.kunrenDB.Collection("words")
}

// StoreWord stores Work if it does not exist or if lemma fingerprint has changed
func (mongo *Mongo) StoreWord(word d.Word) (d.Word, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	// ## find
	fRes := mongo.wordsCollection().FindOne(ctx, bson.D{mp.E{Key: "key", Value: word.Key}})
	fResErr := fRes.Err()

	if fResErr != nil && fResErr != mlib.ErrNoDocuments {
		return d.Word{}, fRes.Err()
	}

	// ## calc hash
	h, hashErr := word.Lemma.Hash()
	if hashErr != nil {
		return d.Word{}, hashErr
	}

	// # if found 1
	if fResErr == nil {
		// ## decode
		var loadedWord d.Word
		err := fRes.Decode(&loadedWord)
		if err != nil {
			return d.Word{}, err
		}

		// ## return ID if hashes match
		if h == loadedWord.LemmaHash {
			return loadedWord, nil
		}

		// ## update
		// update hash
		word.LemmaHash = h
		word.ID = loadedWord.ID
		uRes, err := mongo.wordsCollection().ReplaceOne(ctx, bson.M{"_id": word.ID}, word)
		if err != nil {
			return d.Word{}, err
		}
		if uRes.MatchedCount != 1 {
			return d.Word{}, fmt.Errorf("Nothing was updated for word %v", word.ID)
		}

		fmt.Printf("Updated id %v, old hash %v != new hash %v", word.ID, loadedWord.LemmaHash, h)
		return loadedWord, nil
	}

	// ## otherwise insert
	word.ID = mp.NewObjectID()
	word.LemmaHash = h
	iRes, err := mongo.wordsCollection().InsertOne(ctx, word)
	if err != nil {
		return d.Word{}, err
	}

	fmt.Printf("New word with id %v created", iRes.InsertedID.(mp.ObjectID))
	return word, nil
}

// LoadWord bla
func (mongo *Mongo) LoadWord(id mp.ObjectID) (d.Word, error) {
	var w d.Word
	err := mongo.load(mongo.wordsCollection(), id, &w)
	if err != nil {
		return d.Word{}, err
	}

	return w, nil
}

// DeleteWord bla
func (mongo *Mongo) DeleteWord(id mp.ObjectID) error {
	return mongo.delete(mongo.wordsCollection(), id)
}

// ListWords lists all words in the dict
func (mongo *Mongo) ListWords() ([]d.Word, error) {
	return nil, nil
}
