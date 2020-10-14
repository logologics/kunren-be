package mongo

import (
	"fmt"

	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"

	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

const timeout = 100 * time.Second

// Mongo is a implementation of Repos using MongoDB backend
type Mongo struct {
	kunrenDB *mlib.Database
	client   *mlib.Client
	timeout  time.Duration
}

// Connect creates a new Mongo Repo
func Connect(config *d.Config) (r.Repo, error) {
	client, err := mlib.NewClient(mopt.Client().ApplyURI(config.DB.URL))
	if err != nil {
		return &Mongo{}, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err = client.Connect(ctx)
	if err != nil {
		return &Mongo{}, nil
	}

	kunrenDb := client.Database("kunren")

	m := &Mongo{client: client, kunrenDB: kunrenDb, timeout: timeout}
	err = m.initDB()
	if err != nil {
		return &Mongo{}, err
	}

	return m, nil
}

// Disconnect disconnects
func (m *Mongo) Disconnect() error {
	ctx, _ := context.WithTimeout(context.Background(), m.timeout)
	return m.client.Disconnect(ctx)
}

func hasIndexes(ctx context.Context, iv mlib.IndexView) (bool, error) {
	opts := mopt.ListIndexes().SetMaxTime(2 * time.Second)
	cursor, err := iv.List(ctx, opts)
	if err != nil {
		return false, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return len(results) > 0, nil
}

func (m *Mongo) initDB() error {
	// create indexes
	ctx, _ := context.WithTimeout(context.Background(), m.timeout)
	err := createUserIndexes(ctx, m.kunrenDB)

	return err
}

func (m *Mongo) delete(collection *mlib.Collection, id mp.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	// ## find
	dRes, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if dRes.DeletedCount == 1 {
		log.Info(fmt.Sprintf("Vocab with id %v deleted", id))
	}

	return nil
}
func (m *Mongo) load(collection *mlib.Collection, id mp.ObjectID, target interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	// ## find
	fRes := collection.FindOne(ctx, bson.M{"_id": id})
	fResErr := fRes.Err()

	if fResErr != nil && fResErr != mlib.ErrNoDocuments {
		return fRes.Err()
	}

	// # if not found 1
	if fResErr != nil {
		return fmt.Errorf("Could not find vocab with id  %v", id)
	}

	// ## decode
	err := fRes.Decode(target)
	if err != nil {
		return err
	}

	return nil

}
/*
func Paginate(collection *mongo.Collection, startValue objectid.ObjectID, nPerPage int64) ([]bson.Document, *bson.Value, error) {

	// Query range filter using the default indexed _id field. 
	filter := bson.VC.DocumentFromElements(
			bson.EC.SubDocumentFromElements(
					"_id",
					bson.EC.ObjectID("$gt", startValue),
			),
	)

	var opts []findopt.Find
	opts = append(opts, findopt.Sort(bson.NewDocument(bson.EC.Int32("_id", -1))))
	opts = append(opts, findopt.Limit(nPerPage))

	cursor, _ := collection.Find(context.Background(), filter, opts...)

	var lastValue *bson.Value
	var results []bson.Document
	for cursor.Next(context.Background()) {
			elem := bson.NewDocument()
			err := cursor.Decode(elem)
			if err != nil {
					return results, lastValue, err
			}
			results = append(results, *elem)
			lastValue = elem.Lookup("_id")
	}

	return results, lastValue, nil
}
*/