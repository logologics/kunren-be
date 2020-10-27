package repo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
)

type SortType string

const SortBySeen SortType = "seen"
const SortByConfidence SortType = "conf"

type Repo interface {
	Disconnect() error
	Ready() bool
	WordRepo
	VocabRepo
	UserRepo
}

// WordRepo is used to interact with word storage
type WordRepo interface {
	StoreWord(d.Word) (d.Word, error)
	LoadWord(mp.ObjectID) (d.Word, error)
	DeleteWord(mp.ObjectID) error
	ListWords() ([]d.Word, error)
}

// UserRepo is used to interact with user storage
type UserRepo interface {
	StoreUser(d.User) (d.User, error)
	LoadUser(mp.ObjectID) (d.User, error)
	UpdateUser(d.User) error
	DeleteUser(mp.ObjectID) error
}

// VocabRepo is used to interact with vocab storage
type VocabRepo interface {
	StoreVocab(d.Vocab, bool) (d.Vocab, error)
	LoadVocab(mp.ObjectID) (d.Vocab, error)
	DeleteVocab(mp.ObjectID) error
	ListVocabs(page int, pageSize int, srt []d.Sorting, u d.User) (d.VocabPage, error)
	FindVocab(u d.User, lang d.Language, key string) (d.Vocab, error)
}
