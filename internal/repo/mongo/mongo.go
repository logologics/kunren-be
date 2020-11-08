package mongo

import (
	"fmt"
	"sync"

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
	ready    bool
	mtx      *sync.Mutex
}

func (m *Mongo) setReady() {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.ready = true
	log.Info("Mongo is ready")
}

func (m *Mongo) Ready() bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	return m.ready
}

// Connect creates a new Mongo Repo
func Connect(config *d.Config) (r.Repo, error) {
	client, err := mlib.NewClient(
		mopt.Client().ApplyURI(config.DB.URL).SetServerSelectionTimeout(timeout),
	)
	if err != nil {
		return &Mongo{}, nil
	}

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err = client.Connect(ctx)
	kunrenDb := client.Database("kunren")

	m := &Mongo{client: client, kunrenDB: kunrenDb, timeout: timeout, mtx: &sync.Mutex{}}
	go m.initDB()

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

func (m *Mongo) initDB() {
	for !m.ready {
		log.Info("Trying to create mongo indexes")
		err := m.createIndexes()
		if err == nil {
			m.setReady()
		}
	}
}

func (m *Mongo) createIndexes() error {
	// create indexes
	ctx, _ := context.WithTimeout(context.Background(), m.timeout)
	err := createUserIndexes(ctx, m.kunrenDB)
	if err != nil {
		return err
	}

	err = createWordsIndexes(ctx, m.kunrenDB)
	if err != nil {
		return err
	}

	err = createVocabsIndexes(ctx, m.kunrenDB)
	if err != nil {
		return err
	}

	return nil
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
		log.Info(fmt.Sprintf("Object with id %v deleted", id))
	}

	return nil
}


func (m *Mongo) loadOne(collection *mlib.Collection, id mp.ObjectID, target interface{}) error {
	return m.findOne(collection, bson.M{"_id": id}, target)
}

// FindOne queries a single document with FindOne
func (m *Mongo) findOne(collection *mlib.Collection, query bson.M, target interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	// ## find
	fRes := collection.FindOne(ctx, query)
	fResErr := fRes.Err()

	if fResErr != nil && fResErr != mlib.ErrNoDocuments {
		return fRes.Err()
	}

	// # if not found 1
	if fResErr != nil {
		return fmt.Errorf("Could not find object with query  %v", query)
	}

	// ## decode
	err := fRes.Decode(target)
	if err != nil {
		return err
	}

	return nil

}
