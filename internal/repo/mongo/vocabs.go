package mongo

import (
	"context"
	"fmt"
	"time"

	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
)

func (mongo *Mongo) vocabsCollection() *mlib.Collection {
	return mongo.kunrenDB.Collection("vocabs")
}

// StoreVocab bla
func (mongo *Mongo) StoreVocab(v d.Vocab, inc bool) (d.Vocab, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	// ## find
	fRes := mongo.vocabsCollection().FindOne(ctx, bson.M{
		"wordID": v.WordID,
		"userID": v.UserID,
	})
	fResErr := fRes.Err()

	if fResErr != nil && fResErr != mlib.ErrNoDocuments {
		return d.Vocab{}, fRes.Err()
	}

	// # if found 1
	if fResErr == nil {
		// ## decode
		var loadedVocab d.Vocab
		err := fRes.Decode(&loadedVocab)
		if err != nil {
			return d.Vocab{}, err
		}

		// ## update
		v.ID = loadedVocab.ID
		v.DateSeen = time.Now()
		v.Seen = loadedVocab.Seen

		if inc {
			v.Seen++
		}
		uRes, err := mongo.vocabsCollection().ReplaceOne(ctx, bson.M{"_id": v.ID}, v)
		if err != nil {
			return d.Vocab{}, err
		}
		if uRes.MatchedCount != 1 {
			return d.Vocab{}, fmt.Errorf("Nothing was updated for vocab %v", v.ID)
		}

		return v, nil
	}

	// ## otherwise insert
	v.ID = mp.NewObjectID()
	v.Seen = 1
	iRes, err := mongo.vocabsCollection().InsertOne(ctx, v)
	if err != nil {
		return d.Vocab{}, err
	}

	log.Info(fmt.Sprintf("New vocab with id %v created", iRes.InsertedID))
	return v, nil

}

// LoadVocab blax
func (mongo *Mongo) LoadVocab(id mp.ObjectID) (d.Vocab, error) {
	var v d.Vocab
	err := mongo.load(mongo.wordsCollection(), id, &v)
	if err != nil {
		return d.Vocab{}, err
	}

	return v, nil
}

// DeleteVocab bla
func (mongo *Mongo) DeleteVocab(id mp.ObjectID) error {
	return mongo.delete(mongo.wordsCollection(), id)

}

// ListVocabs lists all Vocabs + words of the user
func (mongo *Mongo) ListVocabs(u d.User, st r.SortType) ([]d.Vocab, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	
		lookupStage := bson.D{
			{"$lookup", bson.D{
				{"from", "words"}, {"localField", "wordID"}, {"foreignField", "_id"}, {"as", "word"}}}}

		unwindStage := bson.D{
			{"$unwind", bson.D{{"path", "$word"}, {"preserveNullAndEmptyArrays", false}}}}
	

	showLoadedCursor, err := mongo.vocabsCollection().Aggregate(
		ctx, mlib.Pipeline{lookupStage, unwindStage})

	if err != nil {
		panic(err)
	}
	var showsLoaded []bson.M
	if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}
	log.Println(showsLoaded)
	return []d.Vocab{}, nil
}
