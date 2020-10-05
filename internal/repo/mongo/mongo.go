package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"

	"context"
	"time"

	log "github.com/sirupsen/logrus"
	mb "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo is a implementation of Repos using MongoDB backend
type Mongo struct {
	kunrenDB *mlib.Database
	client   *mlib.Client
}

// New creates a new Mongo Repo
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

func (m *Mongo) initDB() error {
	// create indexes

	userIdxs := m.kunrenDB.Collection("users").Indexes()
	userIdxmodels := []mlib.IndexModel{
		{
			Keys:    mb.D{primitive.E{Key: "email", Value: 1}},
			Options: mopt.Index().SetName("email_unique").SetUnique(true),
		},
		{
			Keys: mb.D{primitive.E{Key: "name", Value: 1}, primitive.E{Key: "email", Value: 1}},
			Options: mopt.Index().SetName("composite_name_email").SetUnique(true),
		},
	}

	opts := mopt.CreateIndexes().SetMaxTime(2 * time.Second)
	names, err := userIdxs.CreateMany(context.TODO(), userIdxmodels, opts)
	if err != nil {
		return err
	}

	log.Printf("created indexes %v\n", names)

	return nil
}
