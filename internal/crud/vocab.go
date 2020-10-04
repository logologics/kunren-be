package crud

import (
	d "github.com/logologics/kunren-be/internal/domain"
)

func (db *Db) StoreVocab(d.Vocab) error {
	return nil
}

func (db *Db) LoadVocab(id int) (d.Vocab, error) {
	return d.Vocab{}, nil
}

func (db *Db) UpdateVocab(d.Vocab) error {
	return nil
}

func (db *Db) DeleteVocab(id int) error {
	return nil
}
