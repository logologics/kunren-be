package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
)

// StoreVocab bla
func (mongo *Mongo) StoreVocab(d.Vocab) (d.ID, error) {
	return "", nil
}

// LoadVocab blax
func (mongo *Mongo) LoadVocab(d.ID) (d.Vocab, error) {
	return d.Vocab{}, nil
}

// UpdateVocab bla
func (mongo *Mongo) UpdateVocab(d.Vocab) error {
	return nil
}

// UpsertVocab bla
func (mongo *Mongo) UpsertVocab(d.Vocab) (d.ID, error) {
	return "", nil
}

// DeleteVocab bla
func (mongo *Mongo) DeleteVocab(d.ID) error {
	return nil
}

// ListVocab lists all Vocab ofr the user
func (mongo *Mongo) ListVocab(u d.User, st r.SortType) ([]d.Vocab, error){
	return nil, nil
}
