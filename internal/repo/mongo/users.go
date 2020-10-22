package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
	mlib "go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"time"
)

// StoreUser bla
func (mongo *Mongo) StoreUser(d.User) (d.User, error) {
	return d.User{}, nil
}

// LoadUser bla
func (mongo *Mongo) LoadUser(id mp.ObjectID) (d.User, error) {
	return d.User{}, nil
}

//UpdateUser bla
func (mongo *Mongo) UpdateUser(u d.User) error {
	return nil
}

// DeleteUser bla
func (mongo *Mongo) DeleteUser(id mp.ObjectID) error {
	return nil
}

func createUserIndexes(ctx context.Context, db *mlib.Database) error {
	userIdxs := db.Collection("users").Indexes()
	hasIdx, err := hasIndexes(ctx, userIdxs)
	if err != nil {
		return err
	}
	if hasIdx {
		return nil
	}

	userIdxmodels := []mlib.IndexModel{
		{
			Keys:    bson.D{mp.E{Key: "email", Value: 1}},
			Options: mopt.Index().SetName("users_email_unique").SetUnique(true),
		},
		{
			Keys:    bson.D{mp.E{Key: "name", Value: 1}, mp.E{Key: "email", Value: 1}},
			Options: mopt.Index().SetName("users_composite_name_email").SetUnique(true),
		},
	}

	copts := mopt.CreateIndexes().SetMaxTime(2 * time.Second)
	names, err := userIdxs.CreateMany(context.TODO(), userIdxmodels, copts)
	if err != nil {
		return err
	}

	log.Printf("created indexes on users: %v\n", names)

	return nil
}
