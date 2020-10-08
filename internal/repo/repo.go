package repo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
)


type SortType string

const SortBySeen SortType = "seen"
const SortByConfidence SortType = "conf"

type Repo interface{
	Disconnect() error
	WordRepo
	VocabRepo
	UserRepo
}




// WordRepo is used to interact with word storage
type WordRepo interface {
	StoreWord(d.Word) (mp.ObjectID, error)
	LoadWord(mp.ObjectID) (d.Word, error)
	UpdateWord(d.Word) error
	DeleteWord(mp.ObjectID) error
	ListWords() ([]d.Word, error)
}

// UserRepo is used to interact with user storage
type UserRepo interface {
	StoreUser(d.User) (mp.ObjectID, error)
	LoadUser(mp.ObjectID) (d.User, error)
	UpdateUser(d.User) error
	DeleteUser(mp.ObjectID) error
}

// VocabRepo is used to interact with vocab storage
type VocabRepo interface {
	StoreVocab(d.Vocab) (mp.ObjectID, error)
	LoadVocab(mp.ObjectID) (d.Vocab, error)
	UpdateVocab(d.Vocab) error
	UpsertVocab(d.Vocab) (mp.ObjectID, error)
	DeleteVocab(mp.ObjectID) error
	ListVocab(d.User, SortType ) ([]d.Vocab, error)
}


