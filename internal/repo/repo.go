package repo

import (
	d "github.com/logologics/kunren-be/internal/domain"

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
	StoreWord(d.Word) (d.ID, error)
	LoadWord(d.ID) (d.Word, error)
	UpdateWord(d.Word) error
	DeleteWord(d.ID) error
	ListWords() ([]d.Word, error)
}

// UserRepo is used to interact with user storage
type UserRepo interface {
	StoreUser(d.User) (d.ID, error)
	LoadUser(d.ID) (d.User, error)
	UpdateUser(d.User) error
	DeleteUser(d.ID) error
}

// VocabRepo is used to interact with vocab storage
type VocabRepo interface {
	StoreVocab(d.Vocab) (d.ID, error)
	LoadVocab(d.ID) (d.Vocab, error)
	UpdateVocab(d.Vocab) error
	UpsertVocab(d.Vocab) (d.ID, error)
	DeleteVocab(d.ID) error
	ListVocab(d.User, SortType ) ([]d.Vocab, error)
}


