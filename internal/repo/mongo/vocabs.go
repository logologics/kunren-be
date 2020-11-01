package mongo

import (
	"context"
	"fmt"
	"time"

	d "github.com/logologics/kunren-be/internal/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"go4.org/sort"
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
		v.Tags = MergeAndSort(v.Tags, loadedVocab.Tags)

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

		log.Infof("Updated vocab with id %v created\n", v.ID)
		return v, nil
	}

	// ## otherwise insert
	v.ID = mp.NewObjectID()
	v.Seen = 1
	v.DateSeen = time.Now()
	v.DateCreated = v.DateSeen
	iRes, err := mongo.vocabsCollection().InsertOne(ctx, v)
	if err != nil {
		return d.Vocab{}, err
	}

	log.Infof("New vocab with id %v created\n", iRes.InsertedID)
	return v, nil

}

// LoadVocab blax
func (mongo *Mongo) LoadVocab(id mp.ObjectID) (d.Vocab, error) {
	var v d.Vocab
	err := mongo.loadOne(mongo.wordsCollection(), id, &v)
	if err != nil {
		return d.Vocab{}, err
	}

	return v, nil
}

// FindVocab blax
func (mongo *Mongo) FindVocab(u d.User, lang d.Language, key string) (d.Vocab, error) {
	var v d.Vocab
	q := bson.M{"key": key, "userID": u.ID, "language": lang}
	err := mongo.findOne(mongo.vocabsCollection(), q, &v)
	if err != nil {
		return d.Vocab{}, err
	}

	return v, nil
}

// DeleteVocab bla
func (mongo *Mongo) DeleteVocab(id mp.ObjectID) error {
	return mongo.delete(mongo.wordsCollection(), id)

}

func createMatch(uID mp.ObjectID, tags []string) mp.M {
	match := mp.M{
		"userID": uID,
	}
	
	if len(tags) == 0 {
		return match
	}

	bsonTags := mp.A{}
	for _, t := range tags {
		bsonTags = append(bsonTags, t)
	} 
	
	match["tags"] = mp.M{
		"$in": bsonTags,
	}
	return match
}


// ListVocabs lists all Vocabs + words of the user
func (mongo *Mongo) ListVocabs(page int, pageSize int, 
	srt []d.Sorting, u d.User, tags []string) (d.VocabPage, error) {
	
		if page < 0 || pageSize < 1 || page > 250 || pageSize > 50 {
		return d.VocabPage{}, fmt.Errorf("Wrong page/pagesize arg %v/%v", page, pageSize)
	}

	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	skip := page * pageSize

	match := createMatch(u.ID, tags)
	matchStage := bson.D{
		{
			Key: "$match",
			Value: match,
		},
	}

	skipStage := bson.D{
		{
			Key:   "$skip",
			Value: skip,
		},
	}

	limitStage := bson.D{
		{
			Key:   "$limit",
			Value: pageSize,
		},
	}

	var sortStage bson.D
	if len(srt) > 0 {
		sortStage = bson.D{
			{
				Key:   "$sort",
				Value: d.SortingToBson(srt),
			},
		}
	} else {
		sortStage = bson.D{
			{
				Key:   "$sort",
				Value: bson.M{"key": 1},
			},
		}

	}

	lookupStage := bson.D{
		{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: "words"},
				{Key: "localField", Value: "wordID"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "word"},
			},
		},
	}

	unwindStage := bson.D{
		{
			Key: "$unwind",
			Value: bson.D{
				{Key: "path", Value: "$word"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			},
		},
	}

	cur, err := mongo.vocabsCollection().Aggregate(
		ctx, mlib.Pipeline{matchStage, sortStage, skipStage, limitStage, lookupStage, unwindStage},
	)
	if err != nil {
		return d.VocabPage{}, err
	}
	var vocabs []d.VocabListItem
	if err = cur.All(ctx, &vocabs); err != nil {
		return d.VocabPage{}, err
	}

	if vocabs == nil {
		vocabs = []d.VocabListItem{}
	}

	total := mongo.countVocabsForUser(match)
	l := len(vocabs)
	isLast := int64(skip+l) >= total
	isFirst := page == 0
	last := int(total / int64(pageSize))
	if int64(last*pageSize) >= total {
		last--
	}
	return d.VocabPage{
			Vocabs:     vocabs,
			Seq:        page,
			Size:       pageSize,
			Count:      l,
			TotalCount: total,
			IsLast:     isLast,
			IsFirst:    isFirst,
			Last:       last,
		},
		nil
}

// Tags returns a distinct list of all tags the user has created
func (mongo *Mongo) Tags(uID mp.ObjectID) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	matchStage := bson.D{
		{
			Key: "$match",
			Value: mp.M{
				"userID": uID,
			},
		},
	}

	groupStage := bson.D{
		{
			Key: "$group",
			Value: mp.M{
				"_id": "all_tags",
				"tags": mp.M{"$addToSet": "$tags"},
			},
		},
	}

	cur, err := mongo.vocabsCollection().Aggregate(
		ctx, mlib.Pipeline{matchStage, groupStage},
	)

	if err != nil {
		return []string{}, err
	}
	var tagsRaw []bson.D
	if err = cur.All(ctx, &tagsRaw); err != nil {
		return []string{}, err
	}
	log.Infof("got tags %v", tagsRaw)

	tags := joinTagArrays(tagsRaw)

	log.Infof("got tags %v", tags)
	return tags, nil
}

func joinTagArrays(tagsRaw []bson.D) []string {
	m := make(map[string]int)
	for _, d := range tagsRaw {
		tagsList := d.Map()["tags"].(mp.A)
		for _, tagsIf := range tagsList {
			tags := tagsIf.(mp.A)
			for _, tagIf := range tags {
				tag := tagIf.(string)
				if tag != "" {
					m[tag]=1
				}
			}
		}	
	}

	res := make([]string, len(m))
	i:=0
	for k := range m {
		res[i] = k
		i++
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

// DeleteTag deletes a tag from all documents of the user
func (mongo *Mongo) DeleteTag(uID mp.ObjectID, tag string) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	updRes, err := mongo.vocabsCollection().UpdateMany(ctx, bson.M{"userID": uID}, bson.M{
		"$pull": bson.M{"tags": tag},
	})
	if err != nil {
		return err
	}

	log.Infof("Deleted %v from %v documents", tag, updRes.ModifiedCount)
	return nil
}


func (mongo *Mongo) countVocabsForUser(match mp.M) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.timeout)
	defer cancel()

	cnt, err := mongo.vocabsCollection().CountDocuments(ctx, match)
	if err != nil {
		log.Errorf("Count error %v", err)
		return 0
	}

	return cnt
}

func createVocabsIndexes(ctx context.Context, db *mlib.Database) error {
	vocabsIdxs := db.Collection("vocabs").Indexes()
	hasIdx, err := hasIndexes(ctx, vocabsIdxs)
	if err != nil {
		return err
	}
	if hasIdx {
		return nil
	}

	vocabsIdxmodels := []mlib.IndexModel{
		{
			Keys:    bson.D{mp.E{Key: "key", Value: 1}},
			Options: mopt.Index().SetName("vocabs_key"),
		},
		{
			Keys:    bson.D{mp.E{Key: "tags", Value: 1}},
			Options: mopt.Index().SetName("tags_key"),
		},
		{
			Keys:    bson.D{mp.E{Key: "searchStrings", Value: 1}},
			Options: mopt.Index().SetName("vocabs_search_strings"),
		},
		{
			Keys:    bson.D{mp.E{Key: "wordID", Value: 1}, mp.E{Key: "userID", Value: 1}},
			Options: mopt.Index().SetName("vocabs_composite_user_word").SetUnique(true),
		},
	}

	copts := mopt.CreateIndexes().SetMaxTime(2 * time.Second)
	names, err := vocabsIdxs.CreateMany(context.TODO(), vocabsIdxmodels, copts)
	if err != nil {
		return err
	}

	log.Infof("created indexes on vocabs: %v\n", names)

	return nil
}

func MergeAndSort(a []string, b []string) []string {
	res := make([]string, len(a))
	copy(res, a)
	for _, ele := range b {
		res = appendIfMissing(res, ele)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}
