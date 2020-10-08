package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"

	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mlib "go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo is a implementation of Repos using MongoDB backend
type Mongo struct {
	kunrenDB *mlib.Database
	client   *mlib.Client
}

// Connect creates a new Mongo Repo
func Connect(config *d.Config) (r.Repo, error) {
	client, err := mlib.NewClient(mopt.Client().ApplyURI(config.DB.URL))
	if err != nil {
		return &Mongo{}, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return &Mongo{}, nil
	}

	kunrenDb := client.Database("kunren")

	m := &Mongo{client: client, kunrenDB: kunrenDb}
	err = m.initDB()
	if err != nil {
		return &Mongo{}, err
	}

	return m, nil
}

// Disconnect disconnects
func (m *Mongo) Disconnect() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := createUserIndexes(ctx, m.kunrenDB)

	return err 
}
